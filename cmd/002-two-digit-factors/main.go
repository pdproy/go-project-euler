package main

import "fmt"

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

func main() {
	n1 := 1221
	n2 := 3233
	first, second := twoFactors(n1)
	fmt.Printf("factors of %d are: %d, %d\n", n1, first, second)
	first, second = twoFactors(n2)
	fmt.Printf("factors of %d are: %d, %d\n", n2, first, second)
}
