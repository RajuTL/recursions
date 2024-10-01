package main

import "fmt"

// task is to print 1 to 10 without using loops

func printNos(N int) {
	if N == 0 {
		return
	}
	printNos(N - 1)
	fmt.Println(N, " ")
}

func main() {
	N := 10
	printNos(N)
}
