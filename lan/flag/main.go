package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool
	v bool
)

func init() {
	flag.BoolVar(&h, "h", false, "help of this cmd")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	} else if v {
		ver()
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, "Options: \n")
	flag.PrintDefaults()
}

func ver() {
	fmt.Printf("ver = %+v\n", "1.1.0")
}
