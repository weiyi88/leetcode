package main

import "fmt"

//　　（1） 0^0=0，0^1=1 0异或任何数＝任何数
//　　（2） 1^0=1，1^1=0 1异或任何数－任何数取反
//　　（3） 任何数异或自己＝把自己置0

func singleNumber(nums []int) int {
	// 异或，只要两数不相同，结果都为1
	single := 0
	for _, num := range nums {
		single ^= num
		fmt.Println(single)
	}
	return single
}

func main() {
	nums := []int{2, 2, 1}
	singleNumber(nums)
}
