package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/Beadko/gophercises_link/link"
)

var (
	ex = flag.String("example", "ex1.html", "Example file")
)

func main() {
	flag.Parse()
	f, err := os.ReadFile(*ex)
	if err != nil {
		fmt.Printf("Could not open file %s:%v", *ex, err)
	}
	links, err := link.Parse(bytes.NewReader(f))
	if err != nil {
		fmt.Printf("Could not parse the links from HTML %s:%v", *ex, err)
		return
	}
	fmt.Printf("%+v\n", links)
}
