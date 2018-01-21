#### Environment Flag Struct Configuration

*Designed to simplify configuration option mgmt application start up*

Enhanced to include assigning and parsing values from the struct tags,
environment variables and command line flags

- Types are inferred from structure member types
- App prefix inferred from struct name
- Value assignment priority is from top to bottom
  - struct tag
  - environment
  - flag
- Names 
  - environment variable: recursive struct member name with upper case
    camel case underscore separated
  - flag recursive struct member name with lower case camel case
    hyphenated separated
- Flags override environment variables, which in turn override struct
  tags, which override the type default
- Added preprint help text by setting cfg.HelpText before Parse or Process
- json acts as an alias for struct tag name if name tag missing
  - `name:"abc-def"` 
  - `json:"abc-def"`
- If both name and json tags are missing downcase and hyphenated
  struct name is used

```
type S struct {
  i int              `name:"AyeAye"`                      // env name S_AYE_AYE       flag -aye-aye
  f float64          `default:"2.71728"`                  // env name S_F             flag -f
  M map[int]float64 `name:"Map" default:"e:2.71,pi:3.14"` // env name APP_MAP           flag -map
  A []string         `default:"a,b,c"`                    // env name S_A             flag -a
  outer struct {
    i int                                                 // env name S_OUTER_I       flag -outer-i
    inner struct {
       i int                                              // env name S_OUTER_INNER_I flag -outer-inner-i
    }
  }
}
```

- Exporting APP_OVERRIDE_PREFIX set to a value will override the
  structure name for the environment variable prefix e.g. 

```
export APP_OVERRIDE_PREFIX=APP
type S struct {
  i int             `name:"AyeAye"`                       // env name APP_AYE_AYE       flag -aye-aye
  f float64         `default:"2.71728"`                   // env name APP_F             flag -f
  M map[int]float64 `name:"Map" default:"e:2.71,pi:3.14"` // env name APP_MAP           flag -map
  A []string        `default:"a,b,c"`                     // env name APP_A             flag -a
  outer struct {
    i int                                                 // env name APP_OUTER_I       flag -outer-i
    inner struct {
       i int                                              // env name APP_OUTER_INNER_I flag -outer-inner-i
    }
  }
}

```

- The default value can be set in the struct tag with default:"..."
- An environment variable value will replace the tag default if set
- An environment var names come from the member name are prefixed with
  the app prefix
- The flag name derives from the member name, recursive sub struct
  members use the inner struct names as prefixes for both environment
  and flags


```
type S struct {
  i int         
  f float64
  M map[int]float64 `name:"Map" default:"k1:v1,k2:v2,..."`
}

```
- map string representation in tag `default:"k1:v1,k2:v2,..."`
- map string representation in env variable "k1:v1,k2:v2,..."
- slice string

The example/types.go has a sample structure and sample environment
variable bash file which can be tested and run with the following
commands


```
    go get github.com/davidwalter0/go-cfg

    cd ${GOPATH}/src/github.com/davidwalter0/go-cfg/example
    . environment ; go run main.go types.go

```

The current implementation dumps the initialized structure using the
sourced environment variables in the file `environment` merged with
the text from struct tags like: `default:"text"`

The application environment variable prefix is passed to
initialization with the structure to initialize.

The prefix allows override of the environment variables prefix.

Each environment variable is prefixed with the prefix like so

For prefix app and struct member CC the env var will be uppercased and
underscored: `APP_CC`

The command line flag it will be lower cased and hyphenated: `-cc`

But CaC will be `APP_CA_C` and the flag will be -ca-c

Dynamically altering the prefix for an app instance can be done with
by setting the environment variable for the application prefix

Add code similar to the following before calling

`cfg.Initialize(prefix,&structure)`

```
	var prefix = os.Getenv("APP_OVERRIDE_PREFIX")
	if len(prefix) == 0 {
		prefix = "myapp"
	}

    cfg.Initialize(prefix,&structure)
```

Set the appropriate environment variable and run

```
    export APP_OVERRIDE_PREFIX=ANOTHER_PREFIX
```

---

Tags parsed

- name:"nameOverride" which will override existing env variable name
  for this member with the camel case underscore insertion to separate
  words and hyphenation of the lower case of the camel case words
  - environment variable
    `export APP_NAME_OVERRIDE="Tom,Jerry"`
  - flag name
    `-name-override`

- usage:"help text for variable"
- short:"abbrev"
- name:"nameoverride"
- default:"initial value(s) for var conforming to language rules"
  - defaults for slices are comma separated lists
  - defaults for maps are comma separated lists of key:value colon
    separated pairs

Partial example from example/types.go

```
    // Specification example test struct
    type Specification struct {
        Debug bool `name:"Debug" short:"d" default:"false" usage:"enable debug mode"`

```

---

#### Basic use


To use:

```
go get github.com/davidwalter0/go-cfg

```

To use import, create struct info object and call parse.

```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/davidwalter0/go-cfg"
)

type myApp struct {
	I      int `default:"-1"`
	Nested struct {
		Y float64
	}
}

func main() {
	var myapp myApp

	var sti *cfg.StructInfo = &cfg.StructInfo{
		StructPtr: &myapp,
	}

	if err := sti.Parse(); err != nil { // parse tags, environment, flags
		fmt.Errorf("%v\n", err)
	}
	fmt.Printf("%v %T\n", myapp, myapp)
	jsonText, _ := json.MarshalIndent(&myapp, "", "  ")
	fmt.Printf("\n%v\n", string(jsonText))
}

```

Will create a myapp struct with the value -1 in the member variable I,
setting an environment variable `export MYAPP_I=42` will override the
default and the use of the flag `go run example/simple.go -i 42` will
override the environment variable.

Similarly `export MYAPP_NESTED_Y=3.14` will override the value nested
struct value, or `go run example/simple.go --nested-y 3.14` would do
the same

Run with defaults

```
go run example/simple.go
```

Would output

```
{-1 {0}} main.myApp

{
  "I": -1,
  "Nested": {
    "Y": 0
  }
}

```

Running example
```
MYAPP_I=42 go run example/simple.go --nested-y 3.14
```

Would output

```
{42 {3.14}} main.myApp

{
  "I": 42,
  "Nested": {
    "Y": 3.14
  }
}
```


---

A more detailed example is in the example directory and can be run with

```
. example/environment;
go run example/main.go example/myapp.go --user Who \
   -outer-inner-msf  "ξ:1,ρ:0.01,φ:1.2" -map "ξ:1,ρ:0.01,φ:1.2" \
   --name-override "user-a,user-b,user-c" -ca-c x --cc abc -a-b-c-d-e 0xabcde
```

Output will resemble...

```

{
  "Debug": false,
  "Port": 8080,
  "CaC": "x",
  "CC": "abc",
  "User": "",
  "UserName": "",
  "Users": [
    "user-a",
    "user-b",
    "user-c"
  ],
  "UserArray": [
    "x",
    "y",
    "z",
    "0",
    "1"
  ],
  "IntArray": [
    0,
    1,
    2,
    3,
    4
  ],
  "Rate": 2.71828,
  "RateOfTravel": 3.14,
  "Timeout": 2592063000000000,
  "Timeout2": 2592063000000000,
  "Int8": 127,
  "Nint8": -128,
  "Uint8": 255,
  "Int16": 32767,
  "Nint16": -32768,
  "Uint16": 65535,
  "Int32": 1048576,
  "Nint32": -1232,
  "Uint32": 255,
  "ColorCodes": {
    "black": 0,
    "blue": 0,
    "green": 0,
    "red": 0,
    "white": 4095
  },
  "Map": {
    "ξ": 1,
    "ρ": 0.01,
    "φ": 1.2
  },
  "Outer": {
    "I": 42,
    "F": 3.1415926,
    "Msi": {
      "black": 1,
      "blue": 3,
      "green": 2,
      "red": 0,
      "white": 0
    },
    "Inner": {
      "I": 42,
      "F": 3.1415926,
      "Msi": {
        "black": 1,
        "blue": 3,
        "green": 2,
        "red": 0,
        "white": 0
      },
      "Msf": {
        "ξ": 1,
        "ρ": 0.01,
        "φ": 1.2
      }
    }
  },
  "A": {
    "B": {
      "C": {
        "D": {
          "E": 703710
        }
      }
    }
  }
}
```


---
*List & Map Formats*

The flag implementation for slice is limited to comma separated text
conversion to a slice of strings

- "a,b,c" -> []string -> ["a","b","c"]

The map implementation is works with the colon separated key:value
pairs, and commas separating multiple keys:value pairs like
`key1:value1,key2:value2` environment variables.

---
*Acknowledgement*

Initial work was based on the idea / implementation of the envconfig
package in golang; but after refactoring, code generation and import
modification of the google flag package now doesn't share much code
with that work

`github.com/kelseyhightower/envconfig` package

---

Bugs:

Reasonable recursive depth unknown
