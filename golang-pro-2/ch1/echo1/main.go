// Echo1 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func main() {

	//  first version

	// s, sep := "", ""
	// for _, v := range os.Args[1:] {
	// 	s += sep + v
	// 	sep = " "
	// }
	// fmt.Println(s)

	// second version
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// third version
	// fmt.Println(os.Args[1:])

	// ex 1
	// fmt.Println(os.Args[0:])

	// ex 2
	for i, v := range os.Args[0:] {
		fmt.Println(i)
		fmt.Println(v)
	}

	// ex 3

}
