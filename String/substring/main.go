package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "race a car"

	fmt.Println(longestSubString(gString(s)))
}

func longestSubString(strList string) bool {

	letPit := 0
	rightPit := len(strList) - 1
	result := true

	for {
		if strList[letPit] == strList[rightPit] {
			letPit++
			rightPit--
		} else {
			// 不同跳出
			fmt.Printf("%c : %c ", strList[letPit], strList[rightPit])
			fmt.Println()
			result = false
			break
		}

		if letPit > rightPit {
			break
		}
	}
	return result
}

// 判断是否小写字母
func IsString(strI byte) bool {
	if strI >= 97 && strI <= 122 {
		return true
	}
	return false
}

// 提出特殊字符
func gString(str string) string {
	str = strings.ToLower(str)
	var gStr string
	for i := 0; i < len(str); i++ {
		if IsString(str[i]) {
			gStr += string(str[i])
		}
	}
	return gStr
}
