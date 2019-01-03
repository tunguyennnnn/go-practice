package main

import "fmt"

func main() {
	first := make(map[string]int)
	second := map[string]int{
		"k1": 100,
		"k2": 200,
	}

	fmt.Println(first)
	fmt.Println(second)
	for key, value := range second {
		fmt.Println(key, value)
	}

	delete(second, "k1")
	fmt.Println(second)
	value, ok := second["doesntexist"]
	if ok {
		fmt.Println("Exists", value)
	} else {
		fmt.Println("doesnt exist")
	}
}
