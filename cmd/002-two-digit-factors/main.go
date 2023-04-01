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

func compute(n int) {
	first, second := twoFactors(n)
	fmt.Printf("factors of %d are: %d, %d\n", n, first, second)
}
func main() {
	compute(1221)
	compute(3233)
	compute(300)
}
