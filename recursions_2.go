package main

import "fmt"

// task is to print 10 to 1 without using loops
func printNos(N int) {
	if N == 0 {
		return
	}
	// fmt.Print(N, " ")
	fmt.Print(N, " ")
	printNos(N - 1)
}

func main() {
	N := 10
	printNos(N)
}
