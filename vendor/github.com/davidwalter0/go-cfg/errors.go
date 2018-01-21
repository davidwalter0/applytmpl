package cfg

import (
	"errors"
)

var ErrInvalidSpecification = errors.New("Specification must be a struct pointer")
var ErrInvalidArgPointerRequired = errors.New("Argument must be a pointer")
var ErrInvalidArgMapParseSpec = errors.New("Map argument requires pairs")
var ErrIgnoreTag = errors.New("This Tag isn't in use")
