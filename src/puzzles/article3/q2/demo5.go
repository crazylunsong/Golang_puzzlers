package main

import (
	"flag"
	"puzzles/article3/q2/lib"
)

var name string

func init() {
	flag.StringVar(&name, "argsX", "xialunsong", "usage")
}

func main() {
	lib.Hello(name)
}
