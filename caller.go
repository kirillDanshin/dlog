package dlog

import "fmt"

// Caller info
type Caller struct {
	File     string
	Line     int
	FuncName string
}

func (c *Caller) String() string {
	return fmt.Sprintf("Called from %s:%d (%s)", c.File, c.Line, c.FuncName)
}
