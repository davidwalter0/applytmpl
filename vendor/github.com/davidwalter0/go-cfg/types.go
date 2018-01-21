package cfg

// StructInfo struct to hold parse information
type StructInfo struct {
	StructPtr    interface{} // structure pointer to use for initialization
	EnvVarPrefix string      // Application prefix environment variable prefix
	Processed    bool
	EmptyPrefix  bool
}

// MemberType struct to hold parse information
type MemberType struct {
	AppName      string
	EnvVarPrefix string // env variable application prefix
	Name         string // from var name if tag name is present replace member name with tag
	KeyName      string // ENV variable name prefix + "_" + name CamelCase -> PRF_CAMEL_CASE
	Default      string // default from tag, empty string for default
	EnvText      string // environment text, empty string for default
	Short        string // short flag name
	Usage        string // description
	FlagName     string // Hyphenated flag name CamelCase -> camel-case
	Value        string // if env use, else if default tag use, else use type's default
	Depth        int
	SubPrefix    string
	Ignore       string
}
