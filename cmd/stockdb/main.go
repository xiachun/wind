// update datebase of specified report date
package main

import (
	"errors"
	"flag"
	"fmt"
	"net/rpc"
	"os"

	"naturebridge-asset.com/stockdb"
)

var flagsetUpdate *flag.FlagSet

var client *rpc.Client

// define some common errors
var (
	ErrTypeAssertFail = errors.New("type assertion fail")
)

func init() {
	flagsetUpdate = flag.NewFlagSet("update", flag.ContinueOnError)
}

func main() {
	defer stockdb.Close()

	if len(os.Args) < 2 {
		fmt.Println("Usage: stockdb desc|update [options]")
		return
	}

	switch os.Args[1] {
	case "desc":
		stockdb.Describe(os.Stdout)
	case "update":
		update(os.Args[2:])
	default:
		fmt.Println("Usage: stockdb desc|update [options]")
	}
}
