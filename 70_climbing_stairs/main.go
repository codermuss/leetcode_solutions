package main

import "fmt"

func main() {
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}

/*

n=3


*/

/*
n=4
1 1 1 1
1 1 2
1 2 1
2 1 1
2 2

n=5 ->8
n=6 ->  n5-> 8 + n4-> 5 --> 13


*/
