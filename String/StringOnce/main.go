package main

import "fmt"

func main() {
	s := "leetcode"
	fmt.Println(StringOnce(s))
}

func StringOnce(str string) int {
	var cnt = [26]int{}

	for _, v := range str {
		cnt[v-'a']++
	}
	for i, v := range str {
		if cnt[v-'a'] == 1 {
			return i
		}
	}
	return -1
}
