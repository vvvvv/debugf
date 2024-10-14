# delog

`delog` is a simple debugging helper that provides a single function, `delog.Printf`, which offers the same ergonomics as `fmt.Printf`.  
The difference is that output is only printed when compiled with the `delog` build tag.  
Compiling without `delog` makes `delog.Printf` a no-op, eliminating any runtime overhead.  

## Usage

```go
import  "github.com/vvvvv/delog"
main(){
    a:="foo"
    b:="bar"
	delog.Printf("Debug info: %v %v", a, b)
}
```

## Basic Debug Logging With Delog

```go
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

When the above code is build with `delog` enabled by running
```bash
go build -tags delog
```
executing the binary produces the following output:
```
* * * * * * * * * * * * * * * * * * * *
* * * * * * DEBUG BUILD * * * * * * *
* * * * * * * * * * * * * * * * * * * *
- Set DELOG_STACKTRACE=ERROR to see stack traces when delog.Printf() encounters an error.
- Set DELOG_STACKTRACE=ALL to see stack traces on every call to delog.Printf().

hello world
17:28:02 [37.542µs] main.go:12: hi from delog 2024-10-13 17:28:02.248507 +0200 CEST m=+0.000057917
```

#### Note
You can disable the debug build message by providing ldflags during build:
```bash
go build -tags delog -ldflags "-X 'github.com/vvvvv/delog.DisableDebugWarning=1'"
```
This message is printed by default to serve as a warning in production environments.  

## Zero Overhead When Building With delog Disabled
When the code from the previous section is build without the `delog` build tag `go build` only `hello world` gets printed.
Moreover calls to `delog.Printf` are entirely removed at compile time, resulting in zero runtime overhead. This allows you to include `delog.Printf` statements in your code without affecting performance in production builds.  
Using Go's internal objdump tool (`go tool objdump <binary>`) shows the only reference to `delog` in the binary is in fact a no-op.  

```go
TEXT github.com/vvvvv/delog.init.func1(SB) /path/to/delog/delog.go
var Printf PrintfFunc = func(_ string, _ ...interface{}) {}
  0x10009c6b0      d65f03c0        RET
  0x10009c6b4      00000000        ?
  0x10009c6b8      00000000        ?
  0x10009c6bc      00000000        ?
```

## Stack Traces 

`delog` can optionally print stack traces when logging errors. This behavior is controlled by the `DELOG_STACKTRACE` environment variable:

- `DELOG_STACKTRACE=ERROR`: Prints a stack trace when `delog.Printf()` encounters an error.
- `DELOG_STACKTRACE=ALL`: Prints a stack trace on every call to `delog.Printf()`.

```go
package main

import (
	"fmt"

	"github.com/vvvvv/delog"
)

func main() {
	err := fmt.Errorf("some error")
	fmt.Println("hello world")
	delog.Printf("an error occurred: %v", err)
}
```

Compiling with `go build -tags delog` and executing the resulting binary  `DELOG_STACKTRACE=ERROR ./example.bin` prints out:

```
hello world
17:43:48 [37.542µs] main.go:12: an error occurred: some error
goroutine 1 [running]:
github.com/vvvvv/delog.writeStack(0x14000121e98)
        /path/to/delog/delog_build.go:78 +0x74
github.com/vvvvv/delog.printf({0x104a42bbc, 0x15}, {0x140000102d0, 0x1, 0x1})
        /path/to/delog/delog_build.go:62 +0x28c
main.main()
/path/to/example02/main.go:12 +0xe8
```

### See Examples
- [example01](examples/example01) uses the delog build tag.
- [example02](examples/example03) demonstrates stack traces
- [example03](examples/example03) creates two builds, one with delog enabled and one without, and creates objdumps to demonstrate zero overhead.
