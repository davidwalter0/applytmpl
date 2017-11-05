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
	text = buffer.Bytes()
	fmt.Print(buffer)
}
