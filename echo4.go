package main

// Pointers are key to using Go's flag package.

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline") // flag.Bool takes three arguments.
var sep = flag.String("s", " ", "separator")           // flag.String takes a name,

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n { // indirectly acces n
		fmt.Println()
	}
}
