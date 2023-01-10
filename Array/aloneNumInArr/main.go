package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	temArr := make(map[int]int)

	for _, v := range nums {
		temArr[v]++
	}

	for _, v := range nums {
		if temArr[v] > 1 {
			return true
		}
	}
	return false
}
