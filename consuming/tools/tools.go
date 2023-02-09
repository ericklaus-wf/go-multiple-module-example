// Package tools is the recommended way to track toolchain dependencies.
// See https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
package tools

import (
	_ "github.com/ericklaus-wf/go-multiple-module-example/cmd/nomodule"
	_ "github.com/ericklaus-wf/go-multiple-module-example/cmd/taggedmodule"
	_ "github.com/ericklaus-wf/go-multiple-module-example/cmd/untaggedmodule"
)
