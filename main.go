package main

import "fmt"

func Solution(A []int) int  {
	var duplicate map[int]int = make(map[int]int)

	for _, i := range A {
		duplicate[i]++
	}

	fmt.Println(duplicate)

	for k, i := range duplicate {
		if i % 2 == 1 {
			return k
		}
	}

	return 0
}
