// +build !appengine

package flag

import "syscall"

var LookupEnv = syscall.Getenv
