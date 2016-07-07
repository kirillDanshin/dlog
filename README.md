# dlog
Simple build-time controlled debug log

# How to use
### Unbuffered
```
package main

import "github.com/kirillDanshin/dlog"

func main() {
	a := []int{2, 4, 8, 16, 32, 64, 128, 256, 512}
	b := "some string"
	dlog.D(a) // D'ump `a`
	dlog.P(b) // P'rint `b`
	dlog.F("%s format", b) // F'ormatted print
	dlog.Ln(b) // print'Ln `b`
}
```

### Buffered
```
package main

import "github.com/kirillDanshin/dlog"

func main() {
	log := dlog.NewBuffered()
	defer log.Release()
	log.D(a) // D'ump `a`
	log.P(b) // P'rint `b`
	log.F("%s format", b) // F'ormatted print
	log.Ln(b) // print'Ln `b`
}
```

# Release
To disable logging in release build just run
```
	go build
```

# Debug
To enable logging in debug build run
```
	go build -tags "debug"
```
