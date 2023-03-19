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

func twoFactors(n int) (int, int) {
	if n < 10 {
		return 0, 0
	}
	var foundone bool
	var first, second int
	for i := 10; i < 100; i++ {
		if n%i == 0 {
			if !foundone {
				foundone = true
				first = i
			} else {
				second = i
				break
			}
		}
	}
	return first, second
}

func largePalindrome() {
	for n := 9801; n >= 100; n-- {
		if IsIntPalindrome(n) {
			if one, two := twoFactors(n); one != 0 && two != 0 {
				fmt.Printf("%d is largest palindrome that is a product of two 2-git numbers\n", n)
				break
			}
		}
	}
}

func main() {
	largePalindrome()
}
