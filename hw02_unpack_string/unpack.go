package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")
var nums = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

func Unpack(s string) (string, error) {
	// Place your code here.
	unpacked := strings.Builder{}
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		num, ok := nums[r[i]]
		if ok && i > 0 {
			prevRune := string(r[i-1])
			isLastRune := i == len(r)-1
			if !isLastRune {
				nextRune := r[i+1]
				_, ok := nums[nextRune]
				if ok {
					return unpacked.String(), ErrInvalidString
				}
			}
			if num == 0 {
				runeArr := []rune(unpacked.String())
				runeArr = runeArr[:len(runeArr)-1]
				unpacked.Reset()
				unpacked.WriteString(string(runeArr))
			}
			for i := 0; i < num-1; i++ {
				unpacked.WriteString(prevRune)
			}
		} else if curRune := r[i]; curRune == '\\' {
			nextRune := r[i+1]
			_, ok := nums[nextRune]
			if !ok && nextRune != '\\' {
				return unpacked.String(), ErrInvalidString
			}
			unpacked.WriteRune(nextRune)
			i++
		} else if ok && i == 0 {
			return unpacked.String(), ErrInvalidString
		} else {
			unpacked.WriteRune(r[i])
		}
	}
	return unpacked.String(), nil
}
