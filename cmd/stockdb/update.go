package main

import (
	"errors"
	"fmt"
	"log"
	"net/rpc"
	"sync"
	"time"

	"naturebridge-asset.com/stockdb"
	"naturebridge-asset.com/util"
	"naturebridge-asset.com/wind"
)

func union(s1, s2 []string) []string {
	m := make(map[string]struct{})
	for _, s := range s1 {
		m[s] = struct{}{}
	}
	for _, s := range s2 {
		m[s] = struct{}{}
	}

	sUnion := make([]string, 0, len(s1)+len(s2))
	for s := range m {
		sUnion = append(sUnion, s)
	}

	return sUnion
}

func findStocksToUpdate(year, quarter int, override bool) ([]string, []time.Time, error) {
	rptDate := stockdb.Period{year, quarter}.Time()

	// Though we may miss the companys listed after the report date but have quit now, for
	// the sake of material and efficiency, we only check the listed companys of the date
	// of report and now.
	allCodesOfTime, err := wind.SectorCons(wind.C_Sec_AllAShare, time.Now())
	if err != nil {
		return nil, nil, err
	}

	allCodesOfReportDate, err := wind.SectorCons(wind.C_Sec_AllAShare, rptDate)
	if err != nil {
		return nil, nil, err
	}

	allCodes := union(allCodesOfTime, allCodesOfReportDate)

	var codesToQueryIssueDate []string

	if override {
		codesToQueryIssueDate = allCodes
	} else {
		for _, code := range allCodes {
			if !stockdb.FsExist(code, year, quarter) {
				codesToQueryIssueDate = append(codesToQueryIssueDate, code)
			}
		}
	}

	if len(codesToQueryIssueDate) == 0 {
		return nil, nil, nil
	}

	// There's no direct way to find whether a stock has the specified financial
	// report. So we check it by query the issue date of specified report, if not
	// exist, wind may return a zero value (year 1899), then it's reasonable to
	// check if the issue date is after the report date
	_, _, _, windData, err := wind.Wss(
		codesToQueryIssueDate,
		[]string{"stm_issuingdate"},
		&wind.Option{
			ReportDate: rptDate.Format("2006-01-02"),
		})

	if err != nil {
		return nil, nil, err
	}

	issueDatesList, ok := windData.([]interface{})
	if !ok {
		return nil, nil, ErrTypeAssertFail
	}

	var codes []string
	var issueDates []time.Time

	for i, t := range issueDatesList {
		issueDate, ok := t.(time.Time)
		if !ok {
			return nil, nil, ErrTypeAssertFail
		}
		if issueDate.After(rptDate) {
			codes = append(codes, codesToQueryIssueDate[i])
			y, m, d := issueDate.Date()
			issueDates = append(issueDates, time.Date(y, m, d, 0, 0, 0, 0, time.Local)) // normalize the date for compare convenience later
		}
	}

	return codes, issueDates, nil
}

func fetchFsFromWind(codes []string, year, quarter int) ([]float64, int, error) {
	const step = 100 // for limition of wind, we cannot query too many stocks one time

	fields, err := wind.Fields(stockdb.AccCodes())
	if err != nil {
		return nil, 0, err
	}

	o := &wind.Option{
		Unit: 1,
		ReportDate: stockdb.Period{
			Year:    year,
			Quarter: quarter}.Time().Format("20060102"),
		ReportType: 1,
	}

	nCodes, nFields := len(codes), len(fields)

	data := make([]float64, 0, nCodes*nFields)

	for nCodesRead := 0; nCodesRead < nCodes; nCodesRead += step {
		var nc int
		if nCodesRead+step > nCodes {
			nc = nCodes - nCodesRead
		} else {
			nc = step
		}

		fmt.Printf("Fetching %d to %d...\n", nCodesRead, nCodesRead+nc-1)
		_, _, _, d, err := wind.Wss(codes[nCodesRead:nCodesRead+nc], fields, o)
		if err != nil {
			return nil, 0, err
		}

		stepData, ok := d.([]float64)
		if !ok {
			return nil, 0, ErrTypeAssertFail
		}

		data = append(data, stepData...)
	}

	return data, nFields, nil
}

func updateFsFromWind(codes []string, issueDates []time.Time, year, quarter int, override bool, remote bool) (int, error) {
	if len(codes) == 0 {
		return 0, nil
	}

	if len(codes) != len(issueDates) {
		return 0, errors.New("the length of codes and issueDates not agree")
	}

	data, nFields, err := fetchFsFromWind(codes, year, quarter)

	nRecUpdatedLocal := 0
	var wg sync.WaitGroup

	for i, code := range codes {
		if err = stockdb.WriteRecord(
			stockdb.FsKey{code, year, quarter},
			issueDates[i],
			data[i*nFields:(i+1)*nFields],
			override,
		); err != nil {
			return nRecUpdatedLocal, err
		}
		nRecUpdatedLocal++

		if !remote {
			continue
		}

		wg.Add(1)
		go func(code string, i int) {
			if err := client.Call(
				"Remote.WriteRec",
				stockdb.ArgRecord{stockdb.FsKey{code, year, quarter},
					issueDates[i],
					data[i*nFields : (i+1)*nFields],
					override},
				nil,
			); err != nil {
				log.Println(err)
			}

			wg.Done()
		}(code, i)
	}

	fmt.Printf("Waiting for remote...")
	wg.Wait()
	fmt.Println("Done")

	return nRecUpdatedLocal, nil
}

// if not exist then append, or replace
func updateInfoFromWind(codes []string, override bool, remote bool) (int, error) {
	var codesToUpdate []string

	if !override {
		for _, code := range codes {
			if !stockdb.InfoExist(code) {
				codesToUpdate = append(codesToUpdate, code)
			}
		}
	} else {
		codesToUpdate = codes
	}

	if len(codesToUpdate) == 0 {
		return 0, nil
	}

	fields := []string{"sec_name", "ipo_date", "briefing", "backdoordate"}
	_, _, _, wData, err := wind.Wss(codesToUpdate, fields, nil)
	if err != nil {
		return 0, err
	}

	data, ok := wData.([]interface{})
	if !ok {
		return 0, ErrTypeAssertFail
	}

	thre := util.Datetime(1980, 1, 1)

	nInfoUpdatedLocal := 0

	nFields := len(fields)

	var wg sync.WaitGroup

	for i, code := range codesToUpdate {
		var name, brief string
		var listData, backdoor time.Time

		name, ok = data[i*nFields].(string)
		listData, ok = data[i*nFields+1].(time.Time)
		brief, ok = data[i*nFields+2].(string)
		backdoor, ok = data[i*nFields+3].(time.Time)

		if !ok {
			return nInfoUpdatedLocal, ErrTypeAssertFail
		}

		if listData.Before(thre) {
			listData = time.Time{}
		}

		if backdoor.Before(thre) {
			backdoor = time.Time{}
		}

		if err = stockdb.WriteInfo(code, name, listData, backdoor, brief, override); err != nil {
			return 0, err
		}
		nInfoUpdatedLocal++

		if !remote {
			continue
		}

		wg.Add(1)
		go func(code, name string, listData, backdoor time.Time, brief string) {
			if err = client.Call("Remote.WriteInfo",
				stockdb.ArgInfo{code, name, listData, backdoor, brief, override},
				nil,
			); err != nil {
				log.Println(err)
			}

			wg.Done()
		}(code, name, listData, backdoor, brief)
	}

	fmt.Printf("Waiting for remote...")
	wg.Wait()
	fmt.Println("Done")

	return nInfoUpdatedLocal, nil
}

func fetchProductsFromWind(codes []string, year int, quarter int) (map[string][]stockdb.Product, error) {
	mapCodeToProducts := make(map[string][]stockdb.Product)

	o := &wind.Option{
		ReportDate: stockdb.Period{year, quarter}.Time().Format("20060102"),
		Unit:       1,
		Order:      0,
	}

	fields := []string{"segment_product_item", "segment_product_sales", "segment_product_cost"}
	nFields := len(fields)

	stillHave := true

	for stillHave {
		stillHave = false

		o.Order++

		_, _, _, wData, err := wind.Wss(codes, fields, o)
		if err != nil {
			return nil, err
		}

		data, ok := wData.([]interface{})
		if !ok {
			return nil, ErrTypeAssertFail
		}

		for i := len(codes) - 1; i >= 0; i-- {

			// if no product record, delete it from next query
			if data[i*nFields] == nil || data[i*nFields+1] == nil || data[i*nFields+2] == nil {
				codes = append(codes[:i], codes[i+1:]...)
				continue
			}

			name, ok := data[i*nFields].(string)
			if !ok {
				return nil, ErrTypeAssertFail
			}

			sales, ok := data[i*nFields+1].(float64)
			if !ok {
				return nil, ErrTypeAssertFail
			}

			cost, ok := data[i*nFields+2].(float64)
			if !ok {
				return nil, ErrTypeAssertFail
			}

			mapCodeToProducts[codes[i]] = append(mapCodeToProducts[codes[i]], stockdb.Product{name, sales, cost})

			stillHave = true
		}
	}

	return mapCodeToProducts, nil
}

func updateProductsFromWind(codes []string, year int, quarter int, remote bool) (int, error) {
	if quarter != 2 && quarter != 4 {
		return 0, nil
	}

	fullYear := quarter == 4

	var codesToUpdate []string
	for _, code := range codes {
		if !stockdb.ProductsExist(stockdb.ProductsKey{code, year, fullYear}) {
			codesToUpdate = append(codesToUpdate, code)
		}
	}

	if len(codesToUpdate) == 0 {
		return 0, nil
	}

	mapCodeToProducts, err := fetchProductsFromWind(codesToUpdate, year, quarter)
	if err != nil {
		return 0, err
	}

	nProductsUpdatedLocal := 0
	var wg sync.WaitGroup

	for code, prod := range mapCodeToProducts {
		key := stockdb.ProductsKey{code, year, fullYear}
		if err := stockdb.WriteProducts(key, prod); err != nil {
			return nProductsUpdatedLocal, err
		}
		nProductsUpdatedLocal++

		if !remote {
			continue
		}

		// update remote
		wg.Add(1)
		go func(key stockdb.ProductsKey, prod []stockdb.Product) {
			if err := client.Call("Remote.WriteProducts", stockdb.ArgProducts{key, prod}, nil); err != nil {
				log.Println(err)
			}
			wg.Done()
		}(key, prod)
	}

	fmt.Printf("Waiting for remote...")
	wg.Wait()
	fmt.Printf("Done\n")

	return nProductsUpdatedLocal, nil
}

// updateFromWind of the specified report period
// return number of records updated local and any error occued
// include both codes upon current time and codes upon report date
func updateFromWind(year, quarter int, override bool, remote bool) {
	codes, issueDates, err := findStocksToUpdate(year, quarter, override)
	if err != nil {
		log.Println(err)
		return
	}

	if len(codes) == 0 {
		fmt.Println("No stock to update.")
		return
	}

	fmt.Printf("Updating stock info...")
	if nInfoUpdatedLocal, err := updateInfoFromWind(codes, override, remote); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%d Done\n", nInfoUpdatedLocal)
	}

	fmt.Printf("Updating breakdown by product...")
	if nProductsUpdatedLocal, err := updateProductsFromWind(codes, year, quarter, remote); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%d Done\n", nProductsUpdatedLocal)
	}

	fmt.Println("Updating financial statement record...")
	if nRecUpdatedLocal, err := updateFsFromWind(codes, issueDates, year, quarter, override, remote); err != nil {
		log.Println(err)
	} else {
		fmt.Printf("%d Done\n", nRecUpdatedLocal)
	}

	return
}

func mostRecentQuarterBeforeToday() (year, quarter int) {
	now := time.Now()

	year, quarter = now.Year(), (int(now.Month())-1)/3
	if quarter == 0 {
		quarter = 4
		year--
	}

	return
}

func parseUpdateArgs(args []string) (year int, quarter int, override bool, remote bool, err error) {
	defaultYear, defaultQuarter := mostRecentQuarterBeforeToday()

	var updateLocalOnly bool

	flagsetUpdate.IntVar(&year, "y", defaultYear, "report year")
	flagsetUpdate.IntVar(&quarter, "q", defaultQuarter, "report quarter")
	flagsetUpdate.BoolVar(&override, "o", false, "true to override existing record")
	flagsetUpdate.BoolVar(&updateLocalOnly, "l", false, "true to update only local")

	err = flagsetUpdate.Parse(args)

	remote = !updateLocalOnly

	return
}

func update(args []string) {
	year, quarter, override, remote, err := parseUpdateArgs(args)
	if err != nil {
		return
	}

	if remote {
		if client, err = rpc.DialHTTP("tcp", "naturebridge-asset.com:1234"); err != nil {
			log.Println(err)
			return
		}
		defer func() {
			client.Call("Remote.Flush", 0, nil)
			client.Close()
		}()
	}

	updateFromWind(year, quarter, override, remote)
}
