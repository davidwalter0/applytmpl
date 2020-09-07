package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"

	"github.com/davidwalter0/applytmpl"
	"github.com/davidwalter0/applytmpl/doc"
)

var err error
var text []byte

var tmpl *template.Template
var buffer = new(bytes.Buffer)

func closer() {
	if err := os.Stdout.Close(); err != nil {
		log.Fatal(err)
	}
}

var help = flag.Bool("help", false, "print usage information")

func version() {
	me := path.Base(os.Args[0])
	for _, arg := range os.Args {
		if arg == "version" {
			fmt.Fprintf(os.Stderr, "Command: %s\nVersion: %s\nBuild:   %s\nCommit:  %s\n", me, Version, Build, Commit)
			os.Exit(0)
		}
	}
	flag.Parse()
	if *help {
		Usage()
	}
}
func Usage() {
	fmt.Printf(`
%s

Usage: %s [--help]
 
Setting the environment variable HELP=(true,1,t,T) will print this
help

%s version 

prints the build information

Override delimiters for mixed template processing arguments by setting
the environment variable "OVERRIDE_TEMPLATE_DELIMS" to a comma
delimited pair of sequences:
e.g.

OVERRIDE_TEMPLATE_DELIMS=">>,<<" 
or 
OVERRIDE_TEMPLATE_DELIMS="{%%,%%}"

Method names exposed for templates
----------------------------------
`, fmt.Sprintf("Command: %s\nVersion: %s\nBuild:   %s\nCommit:  %s\n", path.Base(os.Args[0]), Version, Build, Commit), os.Args[0], path.Base(os.Args[0]))

	names := applytmpl.SortableStrings{}
	for name := range applytmpl.TemplateFunctions {
		names = append(names, name)
	}

	fmt.Printf("%-32s%s\n", "template function name", "function description")
	fmt.Printf("%-32s%s\n", "--------------------------------", "--------------------------------")
	for _, name := range names.Sort() {
		infoArray := strings.Split(doc.Info(applytmpl.TemplateFunctions[name]), "\n")
		info := strings.Join(infoArray, "\n                                ")
		fmt.Printf("%-32s%s\n", name, info)
	}
	os.Exit(0)
}

func main() {
	version()
	if help, ok := os.LookupEnv("HELP"); ok {
		if v, err := strconv.ParseBool(help); err == nil && v {
			Usage()
		}
	}
	ldelim, rdelim := "{{", "}}"
	if delims, ok := os.LookupEnv("OVERRIDE_TEMPLATE_DELIMS"); ok {
		pair := strings.Split(delims, ",")
		if len(pair) > 1 {
			ldelim = strings.TrimSpace(pair[0])
			rdelim = strings.TrimSpace(pair[1])
		}
	}

	defer closer()
	applytmpl.LoadEnvKV()
	// for k, v := range EnvironmentKV {
	// 	fmt.Println(k, v)
	// }
	text, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("%v.\n", err)
		os.Exit(3)
	}
	// lookup variables and process them first
	tmpl = template.New("TemplateApplyEnv").Delims(ldelim, rdelim)
	tmpl = tmpl.Funcs(applytmpl.TemplateFunctions)
	tmpl, err = tmpl.Parse(string(text))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	tmpl = tmpl.Funcs(applytmpl.TemplateFunctions)
	err = tmpl.Execute(buffer, applytmpl.EnvironmentKV)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if len(applytmpl.Errors) > 0 {
		fmt.Fprintln(os.Stderr, strings.Join(applytmpl.Errors, "\n"))
		os.Exit(1)
	}
	// fmt.Fprintln(os.Stderr, EnvironmentKV)
	text := string(buffer.Bytes())
	find := "<no value"
	// case 0: not found == -1, skip block
	// case 1: found >=0, enter block
	if pos := strings.Index(text, find); pos != -1 {
		var start, hold, line int
		for start = strings.Index(text, "\n"); start >= 0 && start < pos && start != -1; line++ {
			hold = start
			var t int
			if t = strings.Index(text[start:], "\n"); t == -1 {
				break
			}
			start = start + 1 + t
		}
		start = hold
		end := strings.Index(text[start:], "\n")

		if end == -1 {
			end = pos + len(find)
		} else {
			end = start + pos + len(find)
		}
		if end > len(text) {
			end = len(text) - 1
		}
		errorText := strings.Replace(text[start:end], "\n", "", -1)
		fmt.Fprintf(os.Stderr, "Error: template parse failure: [<no value>]\nError: Failure context:%d:%d:\n%s\n", line, end, errorText)
		fmt.Print(text)
		os.Exit(1)
	}
	fmt.Print(text)
}
