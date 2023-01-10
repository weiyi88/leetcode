package main

import "fmt"

func main() {

	s := []byte{'h', 'e', 'l', 'l', 'o'}

	for _, v := range reverString(s) {
		fmt.Printf("%c", v)
	}
}

func reverString(str []byte) []byte {
	for i := 0; i < len(str)/2; i++ {
		var tem byte
		tem = str[i]
		str[i] = str[len(str)-i-1]
		str[len(str)-i-1] = tem
	}
	return str
}
