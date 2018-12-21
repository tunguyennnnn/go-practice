package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4}
	fmt.Println(sl)

	sl2 := make([]int, 20) // make slice length 20

	fmt.Println(sl2)
	for i := 0; i < len(sl2); i++ { //set value to slice item
		sl2[i] = i + 1
	}
	fmt.Println(sl2)
	sl2 = nil // empty the slice
	fmt.Println(sl2)

	for i := 0; i < 10; i++ { // expand slice
		sl2 = append(sl2, i+1)
	}
	fmt.Println(sl2)

	sl2 = sl2[len(sl2)/2 : len(sl2)] // get slice second half
	fmt.Println(sl2)

	sl2 = append(sl2[:2], sl2[3:]...) // remove item at index 2
	fmt.Println(sl2)
}
