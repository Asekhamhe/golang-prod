package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// getInput prompts the user for some text, and then reads a line of input from standard input. This line of text is then returned

func getInput() string {
	fmt.Print("Enter a sentence: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func isNotLetter(c rune) bool {
	return !unicode.IsLetter(c)
}

// isPalindromicSentence returns whether or not the given sentence is palindromic. To calculate this, it splits the string into words, then creates a reverse is equal (ignoring case) to the original.

// It also ignores any non-alphabetic characters
func isPalindromicSentence(s string) bool {
	// split into words and remove non-alphabetic characters ino one operation by using FieldFunc and passing in isNotLetter as the function to split on
	w := strings.FieldsFunc(s, isNotLetter)

	// iterate over the words from front and back simultaneously. if we find a word that is not the same as the word at its matching from the back, the sentence is not palindromic.
	l := len(w)
	for i := 0; i < l/2; i++ {
		fw := w[i]     // front word
		bw := w[l-i-1] //back word
		if !strings.EqualFold(fw, bw) {
			return false
		}
	}

	// all the words matched, so the sentence must be palindromic.
	return true
}

func main() {
	// Go doesn't have while loops, but we can use for loop syntax to read into a new variable, check that it's not empty, and read new lines on subsequent iterations.
	for l := getInput(); l != ""; l = getInput() {
		if isPalindromicSentence(l) {
			fmt.Println(".... is palindromic")
		} else {
			fmt.Println("... is not palindromic.")
		}
	}

}
