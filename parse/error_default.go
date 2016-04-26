// +build !debug

package parse

import (
	"fmt"
)

type baseError struct {
	pos
}

func newBaseError(p pos) *baseError {
	return &baseError{p}
}

func (e *baseError) setTree(t *Tree) {}

func (e *baseError) sprintf(format string, a ...interface{}) string {
	res := fmt.Sprintf(format, a...)
	return fmt.Sprintf("parse: %s on line %d, column %d", res, e.Line, e.Offset)
}
