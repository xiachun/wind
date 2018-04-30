// Package wind provide a simple golang encapsulation of Wind DLL for python

// +build windows

package wind

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"naturebridge-asset.com/util"
)

const (
	windPathFile = "WindPy.pth"
	windDll      = "WindPy.dll"
)

var windDllPath = `C:\Wind\Wind.NET.Client\WindNET\x64\WindPy.dll`
var start, stop, isConnected, wsd, wss, wset, freeData *syscall.LazyProc

// Start Wind
func Start() error {
	if Connected() {
		return nil
	}

	timeOut := 120 * 1000

	r, _, _ := start.Call(util.StrPtr(""), util.StrPtr(""), uintptr(timeOut))

	switch int(r) {
	case 0:
		return nil
	case -40520009:
		return fmt.Errorf("WBox lost")
	case -40520008:
		return fmt.Errorf("Timeout")
	case -40520004:
		return fmt.Errorf("Login failed")
	case -40520014:
		return fmt.Errorf("Please logon iWind first")
	default:
		return fmt.Errorf("Unknow error")
	}
}

// Stop Wind
func Stop() {
	stop.Call()
}

// Connected return connection status
func Connected() bool {
	ok, _, _ := isConnected.Call()
	if ok == 0 {
		return false
	}
	return true
}

// Wss retrieve sectional data from wind
func Wss(codes, fields []string, option *Option) (wCodes []string, wFields []string, wTimes []time.Time, wData interface{}, err error) {
	if err := Start(); err != nil {
		log.Fatal(err)
	}

	r, _, _ := wss.Call(
		util.StrPtr(strings.Join(codes, ";")),
		util.StrPtr(strings.Join(fields, ";")),
		util.StrPtr(option.String()))
	defer freeData.Call(r)

	return distill(r)
}

// Wsd retrieve the time series data
func Wsd(codes, fields []string, begintime, endtime time.Time, option *Option) (wCodes []string, wFields []string, wTimes []time.Time, wData interface{}, err error) {
	if err := Start(); err != nil {
		log.Fatal(err)
	}

	r, _, _ := wsd.Call(
		util.StrPtr(strings.Join(codes, ";")),
		util.StrPtr(strings.Join(fields, ";")),
		util.StrPtr(begintime.Format("2006-01-02")),
		util.StrPtr(endtime.Format("2006-01-02")),
		util.StrPtr(option.String()))
	defer freeData.Call(r)

	return distill(r)
}

// Wset retrieve the predefined dataset
func Wset(table string, option *Option) (wCodes []string, wFields []string, wTimes []time.Time, wData interface{}, err error) {
	if err := Start(); err != nil {
		log.Fatal(err)
	}

	r, _, _ := wset.Call(
		util.StrPtr(table),
		util.StrPtr(option.String()))
	defer freeData.Call(r)

	return distill(r)
}

func distill(r uintptr) (wCodes []string, wFields []string, wTimes []time.Time, wData interface{}, err error) {
	rsWind := (*apiout)(unsafe.Pointer(r))
	if rsWind.errorCode != 0 {
		if s, ok := rsWind.data.getData().([]interface{}); ok {
			sErr := ""
			for i := 0; i < len(s); i++ {
				if ss, ok := s[i].(string); ok {
					sErr += ss
				}
			}
			err = fmt.Errorf(sErr)
		} else {
			err = fmt.Errorf("Unknown error!")
		}
		return
	}

	var ok bool
	wCodes, ok = rsWind.codes.getData().([]string)
	wFields, ok = rsWind.fields.getData().([]string)
	wTimes, ok = rsWind.times.getData().([]time.Time)
	wData = rsWind.data.getData()
	if !ok || wData == nil {
		err = fmt.Errorf("Error get data.")
	}

	return
}

func init() {
	usePathFile := false

	p := filepath.Join("", windPathFile) // Try current directory First
	f, err := os.Open(p)
	if err == nil {
		usePathFile = true
	} else { // Try python path
		pythonPath, err := exec.LookPath("python.exe")
		if err == nil {
			p = filepath.Join(filepath.Dir(pythonPath), "Lib", "site-packages", windPathFile)
			f, err = os.Open(p)
			if err == nil {
				usePathFile = true
			}
		}
	}

	if usePathFile {
		defer f.Close()
		s := bufio.NewScanner(f)
		if s.Scan() {
			windDllPath = filepath.Join(s.Text(), windDll)
		}
	}

	w := syscall.NewLazyDLL(windDllPath)
	start = w.NewProc("start")
	stop = w.NewProc("stop")
	isConnected = w.NewProc("isConnectionOK")
	wsd = w.NewProc("wsd")
	wss = w.NewProc("wss")
	wset = w.NewProc("wset")
	freeData = w.NewProc("free_data")

	// if err = Start(); err != nil {
	// 	log.Fatal(err)
	// }
}
