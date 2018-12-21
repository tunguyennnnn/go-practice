package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height float32
	width  float32
}

func main() {
	sl := make([]Person, 0)
	sl = append(sl, Person{"A", 170.12, 30})
	sl = append(sl, Person{"B", 180, 30})
	sl = append(sl, Person{"C", 120.12, 30})
	sl = append(sl, Person{"D", 150.12, 30})
	fmt.Println(sl)
	sort.Slice(sl, func(i, j int) bool { // ASC
		return sl[i].height < sl[j].height
	})
	fmt.Println("Sorting", sl)

	sort.Slice(sl, func(i, j int) bool { // DESC
		return sl[i].height > sl[j].height
	})
	fmt.Println("Sorting", sl)
}
