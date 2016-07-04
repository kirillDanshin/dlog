// +build debug

package dlog

import "log"

// D dumps a value
func D(v ...interface{}) {
	for _, v := range v {
		log.Printf("%+#v", v)
	}
}

// F is a build-time enabled printf
func F(f string, v ...interface{}) {
	log.Printf(f, v...)
}

// P is a build-time enabled print
func P(v ...interface{}) {
	log.Print(v...)
}

// Ln is a build-time enabled println
func Ln(v ...interface{}) {
	log.Println(v...)
}
