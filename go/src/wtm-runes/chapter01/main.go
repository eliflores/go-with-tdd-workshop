package main

import (
	"fmt"
	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}

// Method declaration (receives a param and only one)
func (cn CharName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", cn.Char, cn.Char, cn.Name)
}

func main() {
	fmt.Println("Please provide one or more words to search.")
}

// Function declaration
func report(word string) {
	char := '\u2108'
	name := runenames.Name(char)
	fmt.Printf("%U\t%c\t%s\n", char, char, name)
	count := 1
	fmt.Printf("%d character found", count)
}
