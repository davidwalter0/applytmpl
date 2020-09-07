package cfg

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

var debug bool

// Debug set the debug test var to true
func Debug(args ...bool) {
	if len(args) == 0 {
		debug = true
		return
	}
	debug = args[0]
}

// FieldPtr for the struct field field
type FieldPtr interface{}

// Field holds the parsed struct tag information
type Field struct {
	reflect.StructField
	FieldPtr
	UseFlags   bool
	StructName string // The name of the current owning structure
	Name       string // from var name if tag name is present replace tag name with tag
	KeyName    string // ENV variable name prefix Prf + "_" + name CamelCase -> PRF_CAMEL_CASE
	Default    string // default from tag, empty string for default
	EnvText    string // environment text, empty string for default
	Short      string // short flag name
	Doc        string // description
	FlagName   string // Hyphenated flag name CamelCase -> camel-case
	Value      string // if env use, else if default tag use, else use type's default
	Omit       bool   // obey json:"...,omitempty" or json:"...,omit" or json:"-"
	Required   bool   // set to force field to have a value
	Depth      int    // struct nesting depth
	Ignore     bool   // don't store or load the corresponding Attribute
	Error      error
	Type       string
	Prefix     string
}

// Get a tag from the struct tags
func (field *Field) Get(name string) string {
	text := field.StructField.Tag.Get(name)
	if len(text) > 0 {
		return text
	}
	return ""
}

// SetField from the struct tags, env, or interpolated values
func (field *Field) SetField() {
	defer func() {
		if err := recover(); err != nil {
		}
	}()

	field.SetDefault()
	field.SetName()
	field.SetIgnore()
	field.SetDoc()
	//	field.SetShort()
	field.SetOmit()
	field.SetRequired()
	field.SetKeyName()
	field.SetFlagName()
	field.SetValueFromEnv()
	field.SetType()
	// Must run last if referencing any info from field settings
	if field.UseFlags {
		field.AddFlag()
	}
}

// String formats args as yaml string
func (field *Field) String() string {
	text, err := yaml.Marshal(*field)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	return string(text)
}

// SetOmit read tag omit option and set when enabled, via ,omit
// ,omitempty or '-' the hyphen option
func (field *Field) SetOmit() {
	json := field.Get("json")
	field.Omit = json == "-" || strings.Index(json, ",omitempty") >= 0 || strings.Index(json, ",omit") >= 0
}

// SetDefault read tag default option and save the text
func (field *Field) SetDefault() {
	field.Default = field.Get("default")
	field.Value = field.Default
}

// SetIgnore read tag ignore option and save the text
func (field *Field) SetIgnore() {
	text := field.Get("ignore")
	if v, err := strconv.ParseBool(text); err == nil {
		field.Ignore = v
	}
}

// SetDoc read tag doc option and save the text
func (field *Field) SetDoc() {
	field.Doc = field.Get("doc")
	if len(field.Doc) == 0 {
		field.Doc = field.Get("help")
	}
	if len(field.Doc) == 0 {
		field.Doc = field.Get("usage")
	}
}

// SetShort read tag short option and save the text
func (field *Field) SetShort() {
	field.Short = field.Get("short")
}

// // SetPrefix read tag prefix option and save the text
// func (field *Field) SetPrefix() {
// 	// field.Prefix = field.Get("prefix")
// }

// SetName read tag name option and save the text
func (field *Field) SetName() {
	name := field.Get("json")
	i := strings.Index(name, ",")
	if i > 0 {
		name = name[0:i]
	}
	if len(name) > 0 {
		field.Name = name
	}
}

// SetIgnore read tag required option and save the text
func (field *Field) SetRequired() {
	text := field.Get("required")
	if v, err := strconv.ParseBool(text); err == nil {
		field.Required = v
	}
}

// SetValueFromEnv uses the value from the environment for this
// structure tag replacing the default tag value
func (field *Field) SetValueFromEnv() {
	if len(field.KeyName) > 0 {
		field.EnvText, _ = LookupEnv(field.KeyName)
	}
	if len(field.EnvText) != 0 {
		field.Value = field.EnvText
	}
}

// SetKeyName read tag keyword option and save the text
func (field *Field) SetKeyName() {
	if len(field.KeyName) == 0 {
		panic("len(field.KeyName) == 0")
	}
	field.KeyNameFromCamelCase()
	field.KeyName = strings.Replace(field.KeyName, "-", "_", -1)
}

// SetFlagName read tag keyword option and save the text
func (field *Field) SetFlagName() {
	if len(field.FlagName) == 0 {
		panic("len(field.FlagName) == 0")
	}
	field.FlagName = Capitalize(field.FlagName)
	field.FlagNameFromCamelCase()
}

// SetType read tag keyword option and save the text
func (field *Field) SetType() {
	// field.Type = field.StructField.Name
}
