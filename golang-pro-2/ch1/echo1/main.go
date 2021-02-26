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
	fmt.Println(os.Args[1:])

}
