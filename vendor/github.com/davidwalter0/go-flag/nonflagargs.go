package flag

import (
	"fmt"
	"os"
	"strconv"
)

// NonFlagOSArgs standard flag package ignores args after the first
// non-flag argument. Move non-flag args to the end of the arg list,
// skip is the number of multicall prefix args.
func NonFlagOSArgs(skip int) []string {
	var multicall = []string{}
	pgm := os.Args[0]
	nonFlagArgs := []string{}
	flagArgs := []string{}
	var i = 1
	// pre-read flag '-' and save non-flag and multicall args
	for ; i < len(os.Args); i++ {
		isFlag := len(os.Args[i]) > 1 && os.Args[i][1] == '-'
		if isFlag {
			flagArgs = os.Args[i:]
			break
		}
		if skip > 0 && i <= skip {
			multicall = append(multicall, os.Args[i])
		} else {
			nonFlagArgs = append(nonFlagArgs, os.Args[i])
		}
	}
	// save flag args and post flag args
	if i < len(os.Args)-1 {
		flagArgs = os.Args[i:]
	}
	os.Args = []string{pgm}
	// flagArgs save non-multicall prefix args from os Args
	// move flag args to the start of the args array
	if len(flagArgs) > 0 {
		os.Args = append(os.Args, flagArgs...)
	}
	// expand non-flag args at the end of the args array
	if len(nonFlagArgs) > 0 {
		os.Args = append(os.Args, nonFlagArgs...)
	}
	return multicall
}

var MultiCallPrefixes []string

func WithMultiCallArgs(args int) {
	if len(MultiCallPrefixes) == 0 {
		MultiCallPrefixes = NonFlagOSArgs(args)
	}
}

// // func init() {
var multicallSetup = func() bool {
	var n int
	v, found := LookupEnv("WITH_MULTICALL_ARGS")
	if found {
		i, err := strconv.ParseInt(v, 0, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		n = int(i)
	}
	MultiCallPrefixes = NonFlagOSArgs(n)
	return true
}()
