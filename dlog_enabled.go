// +build debug

package dlog

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
)

var spewInstance = spew.ConfigState{
	Indent: "\t",
}

// D dumps a value
func D(v ...interface{}) {
	for _, v := range v {
		spewInstance.Dump(v)
	}
}

// F is a build-time enabled printf
func F(f string, v ...interface{}) {
	// log.Printf(f, v...)
	spewInstance.Printf(f+"\n", v...)
}

// P is a build-time enabled print
func P(v ...interface{}) {
	log.Print(v...)
	fmt.Println()
}

// Ln is a build-time enabled println
func Ln(v ...interface{}) {
	log.Println(v...)
}
