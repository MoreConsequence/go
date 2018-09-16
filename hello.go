package main

import (
	"flag"
	lib "hello/lib"
	_ "fmt"
)
var name string

func init(){
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
}