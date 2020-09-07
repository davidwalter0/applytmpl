package cfg

import (
	"fmt"
	"os"
	"strings"

	"github.com/davidwalter0/go-flag"
)

var PrefixOverrideText = `
The environment variable CFG_KEY_PREFIX overrides the prefix
of each environment variable to value. The value will be converted to
upper case like ex - EX

For a struct 

type App struct {
   i int
   s string
}

When CFG_KEY_PREFIX=ex, the environment vars will be
EX_I
EX_S

and the flags will be

--ex-i --ex-s

When CFG_KEY_PREFIX="", the environment vars will be
I
S

and the flags will be

--i --s

Without CFG_KEY_PREFIX override

APP_I
APP_S

and the flags will be
--app-i --app-s
`

var helpText string

// HelpText pre write text for help
func HelpText(text string) {
	helpText = text
}

// Usage add usage text for flags/help processing
var Usage = func() {
	var parts = strings.Split(os.Args[0], "/")
	var l = len(parts)
	var Program = parts[l-1]
	fmt.Fprintf(os.Stderr, "\nUsage of %s:\n", Program)
	if len(helpText) > 0 {
		fmt.Fprintf(os.Stderr, "\n%s\n\n", helpText)
	}

	flag.PrintDefaults()
	// fmt.Fprintf(os.Stderr, PrefixOverrideText)
}

var setup = func() bool {
	flag.Usage = Usage
	return true
}()
