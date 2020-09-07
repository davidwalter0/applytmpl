package cfg

import (
	"regexp"
	"strings"
)

// regexp and replace all from
// https://www.golangprograms.com/golang-convert-string-into-snake-case.html
var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchMultipleHyphen = regexp.MustCompile("([-])([-]*)")
var matchMultipleUnderscore = regexp.MustCompile("([_])([_]*)")

// ToCamelCase from either case test. Doesn't attempt to identify
// acronyms for text replacement; e.g. json_map isn't converted to
// JSONMap, instead this capitalizes the first letter of each snake or
// kebob cast subexpression: json-map and json_map wil both convert to
// JsonMap.
func ToCamelCase(str string) string {
	str = strings.ReplaceAll(str, "_", "-")
	strs := strings.Split(str, "-")
	for i, s := range strs {
		strs[i] = Capitalize(s)
	}
	str = strings.Join(strs, "")
	return str
}

// ToLowerSnakeCase lower case after snake casing string splitting
// CamelCase separating with '_' underscores
func ToLowerSnakeCase(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchMultipleUnderscore.ReplaceAllString(snake, "${1}")
	return strings.ToLower(snake)
}

// ToUpperSnakeCase upper case after snake casing string splitting
// CamelCase separating with '_' underscores
func ToUpperSnakeCase(str string) string {
	str = strings.ReplaceAll(str, "-", "_")
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchMultipleUnderscore.ReplaceAllString(snake, "${1}")
	return strings.ToUpper(snake)
}

// ToLowerKebabCase lower kebab (hyphen case) case splitting from
// CamelCase hyphenating and lower casing camel-case
func ToLowerKebabCase(str string) string {
	str = strings.ReplaceAll(str, "_", "-")
	kebab := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	kebab = matchMultipleHyphen.ReplaceAllString(kebab, "${1}")
	return strings.ToLower(kebab)
}

// ToUpperKebabCase upper kebab (hyphen case) case splitting from
// CamelCase hyphenating and upper casing CAMEL-CASE
func ToUpperKebabCase(str string) string {
	str = strings.ReplaceAll(str, "_", "-")
	kebab := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	kebab = matchAllCap.ReplaceAllString(kebab, "${1}-${2}")
	kebab = matchMultipleHyphen.ReplaceAllString(kebab, "${1}")
	return strings.ToUpper(kebab)
}

// KeyNameFromCamelCase split and upper case splitting from CamelCase
// hyphenating and lower casing CAMEL_CASE. If field.Prefix is set,
// prepend to the text for environment variables. Split via camel case
// regular expression.
func (field *Field) KeyNameFromCamelCase() {
	field.KeyName = ToUpperSnakeCase(strings.Replace(field.KeyName, "-", "_", -1))
}

// FlagNameFromCamelCase for flags CamelCase to camel-case
// hyphenated, split on camel case regular expression
func (field *Field) FlagNameFromCamelCase() {
	field.FlagName = ToLowerKebabCase(strings.Replace(field.FlagName, "_", "-", -1))
}

// ToLowerKebabCase for flags CamelCase to camel-case
// hyphenated, split on camel case regular expression
func (field *Field) ToLowerKebabCase() {
	field.FlagName = ToLowerKebabCase(strings.Replace(field.FlagName, "_", "-", -1))
}

// Capitalize text
func Capitalize(text string) string {
	switch len(text) {
	case 0:
	case 1:
		text = strings.ToUpper(text[0:1])
	default:
		text = strings.ToUpper(text[0:1]) + text[1:]
	}
	return text
}

// Downcase text
func Downcase(text string) string {
	if len(text) > 0 {
		text = strings.ToLower(text)
	}
	return text
}
