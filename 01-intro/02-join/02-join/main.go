package main

import (
	"fmt"
	"runtime"
)

// what will be the output

func main() {
	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
	}()

	for !done {
	}
	fmt.Println("finished")
}
