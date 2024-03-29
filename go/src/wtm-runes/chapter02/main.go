package chapter02

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/runenames"
)

type CharName struct {
	Char rune
	Name string
}

func (c CharName) String() string {
	return fmt.Sprintf("%U\t%c\t%s", c.Char, c.Char, c.Name)
}

func scan(start, end rune) []CharName {
	result := []CharName{}
	for r := start; r < end; r++ {
		name := runenames.Name(r)
		if len(name) > 0 && !strings.HasPrefix(name, "<") {
			result = append(result, CharName{r, name})
		}
	}
	return result
}

func filter(sample []CharName, words []string) []CharName {
	result := []CharName{}
	for i, s := range words {
		words[i] = strings.ToUpper(s)
	}
	for _, c := range sample {
		name := strings.Replace(c.Name, "-", " ", -1)
		parts := strings.Fields(name)

		// Version that did not work for multiple words in the query.
		//for _, part := range parts {
		//	if words[0] == part {
		//		result = append(result, CharName{c.Char, c.Name})
		//		break
		//	}
		//}

		if containsAll(parts, words) {
			result = append(result, CharName{c.Char, c.Name})
		}
	}
	return result
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func containsAll(haystack []string, needles []string) bool {
	for _, s := range needles {
		if !contains(haystack, s) {
			return false
		}
	}
	return true
}

func report(words ...string) {
	result := filter(scan(' ', unicode.MaxRune), words)
	for _, c := range result {
		fmt.Println(c)
	}
	count := len(result)
	fmt.Printf("%d character(s) found", count)
}

func main() {
	// Only syntactic way of doing test in go
	if len(os.Args) > 1 {
		report(os.Args[1:]...)
	} else {
		fmt.Println("Please provide one or more words to search.")
	}
}
