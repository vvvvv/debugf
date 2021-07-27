# delog

Yea debuggers are nice but don't fool yourself, we all know you rather just use `fmt.Println("************here")`.  
`delog` provides a convenience function for debug logging which only prints when build with flag `delog`, Any other build produces a noop logger with no overhead. 

### example

```
package main

import (
	"fmt"
	"time"

	"github.com/vvvvv/delog"
)

func main() {
	fmt.Println("hello world")
	delog.Printf("hi from delog %v", time.Now())
}
```

Building it with `-tags delog`
```
$ go build -tags delog
$ main.bin

* * * * * * * * * * * * * * * * * * * *
* * * * * * * DEBUG BUILD * * * * * * *
* * * * * * * * * * * * * * * * * * * *
- Set DELOG_STACKTRACE=ERROR to see stack traces when delog.Printf() encounters an error.
- Set DELOG_STACKTRACE=ALL to see stack traces on every call to delog.Printf().

hello world
23:48:22 [72.89µs] main.go:12: hi from delog 2021-07-27 23:48:22.744363 +0200 CEST m=+0.000166948
```

Building it without tags
```
$ go build 
$ main.bin

hello world
```