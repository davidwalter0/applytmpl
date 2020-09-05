package cfg

import (
	"errors"
	"fmt"
)

var ErrInvalidSpecification = errors.New("specification must be a struct pointer")
var ErrInvalidArgPointerRequired = fmt.Errorf("%s: requires one or more struct pointer arguments", Package)
var ErrInvalidArgMapParseSpec = errors.New("map argument requires pairs")
var ErrIgnoreTag = errors.New("this Tag isn't in use")
