# dlog [![GoDoc](https://godoc.org/github.com/kirillDanshin/dlog?status.svg)](https://godoc.org/github.com/kirillDanshin/dlog) [![Go Report Card](https://goreportcard.com/badge/github.com/kirillDanshin/dlog)](https://goreportcard.com/report/github.com/kirillDanshin/dlog)
Simple build-time controlled debug log

# How to use
### Unbuffered
```go
package main

import "github.com/kirillDanshin/dlog"

func main() {
	a := []int{2, 4, 8, 16, 32, 64, 128, 256, 512}
	b := "some string"
	
	dlog.D(a)		// D'ump `a`
	dlog.P(b)		// P'rint `b`
	dlog.F("%s format", b)	// F'ormatted print
	dlog.Ln(b)		// print'Ln `b`
}
```

### Buffered
```go
package main

import "github.com/kirillDanshin/dlog"

func main() {
	log := dlog.NewBuffered()
	defer log.Release()
	
	log.D(a)		// D'ump `a`
	log.P(b)		// P'rint `b`
	log.F("%s format", b)	// F'ormatted print
	log.Ln(b)		// print'Ln `b`

	dlog.Ln(log) // or fmt.Println("log") etc.
}
```

# Release
To disable logging in release build just run
```bash
	go build
```

# Debug
To enable logging in debug build run
```bash
	go build -tags "debug"
```
