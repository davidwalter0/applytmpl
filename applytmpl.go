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

var env = make(map[string]string)
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

	environ := os.Environ()
	for _, e := range environ {
		parts := strings.SplitN(e, "=", 2)
		k, v := string(parts[0]), string(parts[1])
		env[k] = v
		env[UCamelCase(k)] = v
		env[LCamelCase(k)] = v
	}

	text, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("%v.\n", err)
		os.Exit(3)
	}
	// lookup variables and process them first
	tmpl = template.New("TemplateApplyEnv")
	tmpl = tmpl.Funcs(fmap)
	tmpl, err = tmpl.Parse(string(text))
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	err = tmpl.Execute(buffer, env)
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if len(errors) > 0 {
		fmt.Fprintln(os.Stderr, strings.Join(errors, "\n"))
		os.Exit(1)
	}

	text = buffer.Bytes()
	fmt.Print(buffer)
}
