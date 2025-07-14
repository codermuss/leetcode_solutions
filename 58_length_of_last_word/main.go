package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	lastString := strings.Split(s, " ")
	return len(lastString[len(lastString)-1])
}
