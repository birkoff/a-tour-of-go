package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	// fmt.Printf("Fields are: %q", strings.Fields(s))
	fields := strings.Fields(s)
	wordCount := make(map[string]int)
	for _,v := range fields {
		// fmt.Printf("Current field is: %s", v)
		_, ok := wordCount[v]
		if ok {
			wordCount[v] += 1
		} else {
			wordCount[v] = 1 
		}
	}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}

// PASS
//  f("I am learning Go!") = 
//   map[string]int{"Go!":1, "I":1, "am":1, "learning":1}
// PASS
//  f("The quick brown fox jumped over the lazy dog.") = 
//   map[string]int{"The":1, "brown":1, "dog.":1, "fox":1, "jumped":1, "lazy":1, "over":1, "quick":1, "the":1}
// PASS
//  f("I ate a donut. Then I ate another donut.") = 
//   map[string]int{"I":2, "Then":1, "a":1, "another":1, "ate":2, "donut.":2}
// PASS
//  f("A man a plan a canal panama.") = 
//   map[string]int{"A":1, "a":2, "canal":1, "man":1, "panama.":1, "plan":1}