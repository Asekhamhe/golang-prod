// Echo4 prints its command-line arguments.

package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

// Note: if the returned type is a pointer type, therefore the value is an address
var sep = flag.String("s", " ", "seperator")

func main() {
	flag.Parse()
	// *sep ===> access the value of an address indirectly
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
