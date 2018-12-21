package main

import "fmt"

func testCopy1() { // len(a1) > //len(a2)
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{-1, -2, -3}
	copy(arr1, arr2)
	fmt.Println("Slice 1", arr1) // [-1 -2 -3 4 5] because arr1 copy 3 first item from arr2
}

func testCopy2() { // len(a1) < len(a2)
	arr1 := []int{1, 2, 3}
	arr2 := []int{-1, -2, -3, -4, -5}
	copy(arr1, arr2)
	fmt.Println("Slice 1", arr1) // [-1, -2, -3]
}

func testCopy3() { // copy from array to slice
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := []int{-1, -2, -3, -4, -5, -7}
	copy(arr2, arr1[:]) // convert array arr1 to slice
	fmt.Println("Slice 2", arr2)
}

func testCopy4() {
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := []int{-1, -2, -3, -4, -5, -7}
	copy(arr1[:], arr2)
	fmt.Println("Slice 2", arr2)
	fmt.Println("Array", arr1) // [-1, -2, -3, -4]
}

func main() {
	testCopy1()
	testCopy2()
	testCopy3()
	testCopy4()
}
