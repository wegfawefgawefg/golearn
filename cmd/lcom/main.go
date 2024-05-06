package main

import (
	"fmt"
	"trying-mo/pkg/lcom"
)

func main() {
	src := "(add 2 (subtract \"314\" 2))"

	tokens, err := lcom.Tokenize(src)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}
