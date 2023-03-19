package main

import "fmt"

func IsIntPalindrome(n int) bool {
	sn := fmt.Sprintf("%d", n)
	l := len(sn)
	for i := 0; i < l/2; i++ {
		if sn[i] != sn[l-1-i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsIntPalindrome(1221))
	fmt.Println(IsIntPalindrome(12321))
	fmt.Println(IsIntPalindrome(123421))
}
