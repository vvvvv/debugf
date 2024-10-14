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
