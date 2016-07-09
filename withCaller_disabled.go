// +build !debug

package dlog

// D dumps a value
func (*WithCaller) D(v ...interface{}) {}

// F is a build-time enabled printf
func (*WithCaller) F(f string, v ...interface{}) {}

// P is a build-time enabled print
func (*WithCaller) P(v ...interface{}) {}

// Ln is a build-time enabled println
func (*WithCaller) Ln(v ...interface{}) {}

// GetCaller is a build-time disabled caller determining.
// Returns caller's file, line and func name
func GetCaller(_ ...int) (*Caller, bool) {
	return &Caller{
		File:     CallerUnknown,
		Line:     0,
		FuncName: CallerUnknown,
	}, false
}
