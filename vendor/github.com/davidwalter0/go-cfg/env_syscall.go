// +build !appengine

package cfg

import "syscall"

var lookupEnv = syscall.Getenv
