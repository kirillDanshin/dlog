// +build !debug

package dlog

// D dumps a value
func D(v ...interface{}) {}

// F is a build-time disabled printf
func F(v ...interface{}) {}

// P is a build-time disabled print
func P(v ...interface{}) {}

// Ln is a build-time disabled println
func Ln(v ...interface{}) {}
