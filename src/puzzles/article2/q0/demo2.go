package main

import (
	"flag"
	"fmt"
)

var name string
var s = flag.String("name2", "lunsong2", "123")

func init()  {
	flag.StringVar(&name, "name", "lunsong", "test")
}

func main() {
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
	fmt.Printf("Hello, %s!\n", *s)

}
