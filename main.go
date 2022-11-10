package main

import (
	"fmt"
	"strings"
	"os"
)

func main() {
	words := os.Args[1:]
	word := words[0]
	if word != "" {
		listWords := strings.Split(word, " ")
		fmt.Println(len(listWords))
		return
	}
	fmt.Println(0)
	return
}
