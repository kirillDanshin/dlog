// +build debug

package dlog

import (
	"log"
	"runtime"
	"strings"

	"github.com/kirillDanshin/myutils"
)

// D dumps a value
func (*WithCaller) D(v ...interface{}) {
	c, _ := GetCaller()
	log.Print(c, ": [")
	for _, v := range v {
		spewInstance.Dump(v)
	}
	log.Println("]")
}

// F is a build-time enabled printf
func (*WithCaller) F(f string, v ...interface{}) {
	c, _ := GetCaller()
	// log.Printf(myutils.Concat(c.String(), "[\n\t", f, "\n]"), v...)
	spewInstance.Printf(myutils.Concat(c.String(), "[\n\t", f, "\n]"), v...)
}

// P is a build-time enabled print
func (*WithCaller) P(v ...interface{}) {
	c, _ := GetCaller()
	log.Print(c, ": [\t")
	log.Print(v...)
	log.Println("]")
}

// Ln is a build-time enabled println
func (*WithCaller) Ln(v ...interface{}) {
	c, _ := GetCaller()
	log.Print(c, ": [\t")
	log.Println(v...)
	log.Println("]")
}

// GetCaller returns caller's file, line and func name
func GetCaller(stackBack ...int) (*Caller, bool) {
	sb := 2
	if len(stackBack) > 0 {
		sb = stackBack[0] + 1
	}
	pc, file, line, ok := runtime.Caller(sb)
	if !ok {
		return &Caller{
			File:     CallerUnknown,
			Line:     0,
			FuncName: CallerUnknown,
		}, false
	}

	if li := strings.LastIndex(file, "/"); li > 0 {
		file = file[li+1:]
	}

	return &Caller{
		File:     file,
		Line:     line,
		FuncName: runtime.FuncForPC(pc).Name(),
	}, true
}
