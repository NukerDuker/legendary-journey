package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var nums = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

func Unpack(s string) (string, error) {
	// Place your code here.
	result := strings.Builder{}
	for i, r := range s {
		num, ok := nums[string(r)]
		if ok {
			prevLetter := string(s[i-1])
			for i := 0; i < num-1; i++ {
				result.WriteString(prevLetter)
			}
		} else {
			result.WriteString(string(r))
		}
	}
	return result.String(), nil
}
