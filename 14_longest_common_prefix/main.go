package main

import "fmt"

func main() {

	strs := []string{"flower", "flow", "floight"}
	fmt.Println(longestCommonPrefix(strs))

}

func longestCommonPrefix(strs []string) string {
	p := strs[0]

	for j := 1; j < len(strs); j++ {
		i := 0
		for ; i < len(strs[j]) && i < len(p) && p[i] == strs[j][i]; i++ {
		}
		p = p[:i]
	}
	return p
}
