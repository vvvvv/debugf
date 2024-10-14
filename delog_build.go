// +build delog

package delog

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var (
	DisableDebugWarning string = "0"
)

type stackflag uint8

const (
	printstack stackflag = 1 << iota
	onerror
	all
)

const (
	calldepth = 2
	envPrefix = "DELOG_"
	// timeFormat          = "15:04:05" // zone info not needed as this is only used short debugging sessions
)

func env(name string) string {
	return strings.ToLower(os.Getenv(fmt.Sprintf("%s%s", envPrefix, name)))
}

var (
	stackflags stackflag = 0
	mu         sync.Mutex
	out        = os.Stderr
	// size of buffer for Debugf
	bufSize = 1024

	// size of buffer for stack traces
	// will be added to bufSize in init() if build with -tags "debug stack"
	stackBufSize = 1024
)

// sync pool to manage lifetime of buffers
var bufPool = sync.Pool{
	New: func() interface{} { return make([]byte, bufSize) },
}

func printf(f string, v ...interface{}) {
	var b []byte
	b = bufPool.Get().([]byte)

	b = b[:0]
	hasError(v...)

	formatInfo(&b)
	f += "\n"
	b = append(b, fmt.Sprintf(f, v...)...)

	if stackflags&printstack != 0 {
		if (stackflags&onerror != 0) && hasError(v...) {
			writeStack(&b)
		} else if stackflags&onerror == 0 {
			writeStack(&b)
		}

	}

	mu.Lock()
	defer mu.Unlock()
	out.Write(b)

}

func writeStack(buf *[]byte) {
	b := make([]byte, stackBufSize)
	for {
		n := runtime.Stack(b, false)
		if n < len(b) {
			*buf = append(*buf, b...)
			*buf = append(*buf, '\n')
			return
		}
		b = make([]byte, 2*len(b))
	}
}

func hasError(args ...interface{}) bool {
	for _, arg := range args {
		switch arg.(type) {
		case error:
			return true
		default:
		}
	}
	return false
}

func formatInfo(buf *[]byte) {
	now := time.Now()
	since := now.Sub(timeStart)

	_, file, line, ok := runtime.Caller(calldepth)
	if !ok {
		file = "no_file"
		line = 0
	}

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			file = file[i+1:]
			break
		}
	}

	// time
	h, min, sec := now.Clock()
	itoa(buf, h, 2)
	*buf = append(*buf, ':')
	itoa(buf, min, 2)
	*buf = append(*buf, ':')
	itoa(buf, sec, 2)

	// since
	*buf = append(*buf, " ["...)
	*buf = append(*buf, since.String()...)
	*buf = append(*buf, "] "...)

	// file
	*buf = append(*buf, file...)
	*buf = append(*buf, ':')
	// line
	itoa(buf, line, -1)
	*buf = append(*buf, ": "...)

}

// itoa zero padds ints
// from sdtlib log
func itoa(buf *[]byte, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

var hello = `* * * * * * * * * * * * * * * * * * * *
* * * * * * * DEBUG BUILD * * * * * * *
* * * * * * * * * * * * * * * * * * * *
`
var help = `- Set DELOG_STACKTRACE=ERROR to see stack traces when delog.Printf() encounters an error.
- Set DELOG_STACKTRACE=ALL to see stack traces on every call to delog.Printf().
`

func msg() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Fprint(out, hello)
	fmt.Fprint(out, help)
	fmt.Fprint(out, "\n")
}

var timeStart time.Time

func init() {
	timeStart = time.Now()
	if(DisableDebugWarning == "0"){
		msg()
	}

	Printf = printf

	printStack := env("STACKTRACE")
	if printStack != "" {
		stackflags |= printstack
		bufSize += stackBufSize

		switch printStack {
		case "error":
			stackflags |= onerror
		case "all":
			stackflags |= all
		default:
			stackflags |= onerror

		}
	}
}
