package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more float")
		os.Exit(1)
	}
	arguments := os.Args[1:]
	floatList := convertToFloatList(arguments)
	min, max := findMinMax(floatList)
	fmt.Printf("Min is %f \n", min)
	fmt.Printf("Max is %f \n", max)
}

func convertToFloatList(stringList []string) []float64 {
	var resList []float64
	for _, str := range stringList {
		n, err := strconv.ParseFloat(str, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		resList = append(resList, n)
	}
	return resList
}

func findMinMax(floatList []float64) (float64, float64) {
	min := floatList[0]
	max := floatList[0]
	for _, floatNum := range floatList[1:] {
		if min > floatNum {
			min = floatNum
		}
		if max < floatNum {
			max = floatNum
		}
	}
	return min, max
}
