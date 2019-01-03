package main

import "fmt"

func testPointer(n *int) {
	*n = *n * *n
}

func testPointer2(n int) {
	n = n * n
}

func main() {
	n := 10
	testPointer(&n)
	fmt.Println(n)
	testPointer2(n)
	fmt.Println(n)
}
