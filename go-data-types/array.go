package main

import "fmt"

func main() {
	arr := [4]int{1, 2, 3, 4}
	for number, index := range arr {
		fmt.Printf("Number %d at index %d \n", number, index)
	}

	fmt.Println("2d array ______________________")
	twodArr := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	for i := 0; i < len(twodArr); i++ {
		for j := 0; j < len(twodArr[i]); j++ {
			fmt.Printf("Number at row %d column %d is %d\n", i, j, twodArr[i][j])
		}
	}
}
