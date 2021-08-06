package main

import (
	"flag"
	"fmt"
)

var name string

func init()  {
	flag.StringVar(&name, "name", "lunsong", "test")
}

func main() {
	s:=1
	fmt.Println(s)
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)

}
