package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxSonNum(nums))
}

func maxSonNum(arr []int) int {
	temMax := 0
	temp := 0
	for k, v := range arr {
		if v < 0 {
			continue
		}
		for i := k + 1; i < len(arr); i++ {
			if arr[i] < 0 {
				continue
			}
			temp = arr[i] * v
			temMax = Max(temp, temMax)
		}
	}
	return temMax
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b

}
