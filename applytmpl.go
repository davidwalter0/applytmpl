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

// errors list of problems during template processing
var errors []string

// Env lookup name return value
func Env(name string) (text string) {
	text = os.Getenv(name)
	if len(text) == 0 {
		if len(errors) == 0 {
			errors = append(errors, fmt.Sprintf("Template Processing Error"))
		}
		errors = append(errors, fmt.Sprintf("env var unset: %s", name))
	}
	return
}

var fmap = template.FuncMap{
	"env": Env,
}

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
	}

	text, err = ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Printf("%v.\n", err)
		os.Exit(3)
	}

	tmpl = template.New("TemplateApplyString")
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
	fmt.Print(buffer)
}
