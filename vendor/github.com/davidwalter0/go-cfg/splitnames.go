package cfg

import (
	"log"
	"regexp"
	"strings"
)

var regExpr = regexp.MustCompile("([^A-Z]+|[A-Z][^A-Z]+|[A-Z]+)")

func init() {
	if false {
		log.Println("")
	}
}

// UnderScoreCamelCaseWords split on CamelCase words
func (member *MemberType) UnderScoreCamelCaseWords() {
	words := regExpr.FindAllStringSubmatch(member.Name, -1)
	if len(words) > 0 {
		var envSepName []string
		for _, words := range words {
			envSepName = append(envSepName, words[0])
		}
		member.KeyName = strings.Join(envSepName, "_")
		if len(member.EnvVarPrefix) > 0 {
			member.KeyName = strings.ToUpper(member.EnvVarPrefix + "_" + member.KeyName)
		} else {
			member.KeyName = strings.ToUpper(member.KeyName)
		}
	}
}

// HyphenateCamelCaseWords converts camel case name string and
// hyphenates words for flags between words
func (member *MemberType) HyphenateCamelCaseWords() {
	prefix := strings.Replace(member.EnvVarPrefix, member.AppName, "", 1)

	if len(prefix) > 0 && (prefix[0] == '-' || prefix[0] == '_') {
		prefix = prefix[1:]
	}

	if len(prefix) > 0 {
		member.FlagName = prefix + "-" + Capitalize(member.Name)
	}

	member.FlagName = strings.Replace(member.FlagName, "_", "-", -1)
	words := regExpr.FindAllStringSubmatch(member.FlagName, -1)
	if len(words) > 0 {
		var name []string
		for _, words := range words {
			name = append(name, strings.ToLower(words[0]))
		}
		member.FlagName = strings.Join(name, "-")
	}
	for n := strings.Index(member.FlagName, "--"); n > 0; n = strings.Index(member.FlagName, "--") {
		member.FlagName = strings.Replace(member.FlagName, "--", "-", -1)
	}
}
