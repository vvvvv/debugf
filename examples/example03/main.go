package main

import (
	"fmt"

	"github.com/vvvvv/delog"
)

func main() {
	err := fmt.Errorf("some error")
	fmt.Println("hello world")
	// this doesn't produce any output because it wasn't build with -tag delog
	delog.Printf("an error occurred: %v", err)
}
