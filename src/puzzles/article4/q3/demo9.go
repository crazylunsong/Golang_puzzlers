package main

import (
	"fmt"
	"sync"
)

func main() {
	var err error
	var b string
	b ="a"
	fmt.Printf(b)
	b = "c"
	fmt.Println(b)
	var a sync.WaitGroup
	fmt.Println(err)
	fmt.Println(a)
}
