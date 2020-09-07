package cfg

import (
	"fmt"
	"strings"

	"github.com/davidwalter0/go-flag"
)

var allFlagNames = make(map[string]bool)
var announceDuplicates bool

// AddFlag if not created and warn if it is a duplicate
func (field *Field) AddFlag() {
	if _, ok := allFlagNames[field.FlagName]; !ok {
		var usage string
		if len(field.Doc) > 0 {
			usage = "usage: " + field.Doc
		}
		isset := len(field.Default) > 0 || len(field.EnvText) > 0
		flag.MakeVar(field.FieldPtr, field.FlagName, field.Default, usage+fmt.Sprintf(" Env %-32s : (%s) (%v)", field.KeyName, field.Name, field.Type), field.Value, field.Required, isset)
	} else {
		if !announceDuplicates {
			fmt.Printf("Duplicate flag(s)/env vars found\n")
			fmt.Println(strings.ToUpper(fmt.Sprintf("%-20s %-20s", "github.com/davidwalter0/go-flag", "env vars")))
			fmt.Println("-----------------------------------------")
			announceDuplicates = true
		}
		fmt.Printf("%-20s %-20s\n", field.FlagName, field.KeyName)
	}
}

// IsSet returns if the flag has been set
func IsSet(name string) bool {
	return flag.IsSet(name)
}

// Ok returns if the flag has been set
func Ok(name string) bool {
	return flag.Ok(name)
}

// Required returns if the flag has been set
func Required(name string) bool {
	return flag.Required(name)
}
