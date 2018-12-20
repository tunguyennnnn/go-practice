package main

import "fmt"

// Defer in go is LIFO
func testDefer() {
	for i := 0; i <= 3; i += 1 {
		fmt.Printf("Before defer of i %d \n", i)
		defer fmt.Printf("value of i is %d \n", i)
	}
}

//
func testDefer2() {
	for i := 0; i <= 3; i += 1 {
		fmt.Printf("Before defer of i %d \n", i)
		defer func() {
			fmt.Printf("value of i is %d \n", i)
		}()
	}
}

func testDefer3() {
	for i := 0; i <= 3; i += 1 {
		fmt.Printf("Before defer of i %d \n", i)
		defer func(i int) {
			fmt.Printf("value of i is %d \n", i)
		}(i)
	}
}

func testPanic2() {
	fmt.Println(" inside test panic 2")
	panic(" panic in panic 2")
	fmt.Println("Exiting panic 2")
}

func testPanic() {
	fmt.Println("inside testPatic()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside testPanic()")
		}
	}()
	fmt.Println("about to call testPanic2()")
	testPanic2()
	fmt.Println("call testPanic2 exited")
	fmt.Println("Exiting testPanic")
}

func main() {
	fmt.Println("Defer 1_______________________")
	testDefer()
	fmt.Println("Defer 2_______________________")
	testDefer2()
	fmt.Println("Defer 3_______________________")
	testDefer3()
	fmt.Println("Test panic _________________")
	testPanic()
	fmt.Println(" end main ")
}
