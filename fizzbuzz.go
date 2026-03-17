// Package fizzbuzz: 3→Fizz, 5→Buzz, 15→FizzBuzz.
package fizzbuzz

import "strconv"

const (
	FizzStr     = "Fizz"
	BuzzStr     = "Buzz"
	FizzBuzzStr = "FizzBuzz"
)

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return FizzBuzzStr
	}
	if n%3 == 0 {
		return FizzStr
	}
	if n%5 == 0 {
		return BuzzStr
	}
	return strconv.Itoa(n)
}
