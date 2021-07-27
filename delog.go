// Package delog is useful for printing debug messages
package delog

// Printf formats exactly like fmt.Printf does but writes to stderr.
// This function will only print if build with `go build -flags delog` and is a noop otherwise.

// go:noinline
var Printf PrintfFunc = func(_ string, _ ...interface{}) {}

type PrintfFunc func(string, ...interface{})
