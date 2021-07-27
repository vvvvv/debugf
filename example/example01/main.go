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
