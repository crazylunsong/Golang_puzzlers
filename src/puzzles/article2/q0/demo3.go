package main

import (
	"flag"
)

var name4 string
var cmdLine = flag.NewFlagSet("question", flag.ExitOnError)


func init() {
	//flag.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	//flag.CommandLine.Usage = func() {
	//	fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
	//	flag.PrintDefaults()
	//}
	//
	////flag.StringVar(&name4, "name", "lunsong", "test")
	//cmdLine.StringVar(&name4, "name", "lunsong", "test")
}

func main() {
	//flag.Parse()
	//cmdLine.Parse(os.Args[0:])
	//fmt.Printf("Hello, %s!\n", name4)

}
