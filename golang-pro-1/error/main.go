package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	m := map[string]int{}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		m[s.Text()]++
	}

	if err := s.Err(); err != nil {
		log.Printf("ERROR: could not read stdin: %s", err)
	}
	for k, v := range m {
		log.Printf("%s => %d", k, v)
	}
}
