package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

var EnvironmentKV map[string]string = make(map[string]string, 0)

var err error
var text []byte

var tmpl *template.Template
var buffer = new(bytes.Buffer)

func closer() {
	if err := os.Stdout.Close(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer closer()
	LoadEnvKV()
	// for k, v := range EnvironmentKV {
	// 	fmt.Println(k, v)
	// }
	text, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("%v.\n", err)
		os.Exit(3)
	}
	// lookup variables and process them first
	tmpl = template.New("TemplateApplyEnv")
	tmpl = tmpl.Funcs(TemplateFunctions)
	tmpl, err = tmpl.Parse(string(text))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	tmpl = tmpl.Funcs(TemplateFunctions)
	err = tmpl.Execute(buffer, EnvironmentKV)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if len(errorStrings) > 0 {
		fmt.Fprintln(os.Stderr, strings.Join(errorStrings, "\n"))
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
