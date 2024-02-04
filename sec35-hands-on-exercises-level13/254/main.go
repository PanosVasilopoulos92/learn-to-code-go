package main

import (
	"fmt"

	"github.com/PanosVasilopoulos92/learn-to-code-go/sec35-hands-on-exercises-level13/254/quote"
	"github.com/PanosVasilopoulos92/learn-to-code-go/sec35-hands-on-exercises-level13/254/word"
)

func main() {
	for k, v := range word.UseCount(quote.QuoteText) {
		fmt.Println(k, ":", v)
	}

	fmt.Println()
	fmt.Println("Number of words:", word.Count(quote.QuoteText))
}
