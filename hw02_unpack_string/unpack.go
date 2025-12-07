package hw02unpackstring

import (
	"errors"
	"strings"
)

var (
	ErrInvalidString = errors.New("invalid string")
	nums             = map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}
)

func Unpack(s string) (string, error) {
	// Place your code here.
	unpacked := &strings.Builder{}
	unpacked, err := iterateStringAndUnpackRunes(s, unpacked)
	if err != nil {
		return unpacked.String(), err
	}

	return unpacked.String(), nil
}

func iterateStringAndUnpackRunes(s string, unpacked *strings.Builder) (*strings.Builder, error) {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		num, ok := nums[r[i]]
		if ok && i > 0 {
			return handleNum(r, unpacked, num, i)
		} else if curRune := r[i]; curRune == '\\' {
			unpacked, err := handleEscapeSlash(r, unpacked, i)
			if err != nil {
				return unpacked, err
			}
			i++
		} else if ok && i == 0 {
			return unpacked, ErrInvalidString
		} else {
			unpacked.WriteRune(r[i])
		}
	}
	return unpacked, nil
}

func handleNum(r []rune, unpacked *strings.Builder, num int, i int) (*strings.Builder, error) {
	prevRune := string(r[i-1])
	isLastRune := i == len(r)-1
	if !isLastRune {
		nextRune := r[i+1]
		_, ok := nums[nextRune]
		if ok {
			return unpacked, ErrInvalidString
		}
	}
	if num == 0 {
		unpacked = removeLastRune(unpacked)
	}
	for i := 0; i < num-1; i++ {
		unpacked.WriteString(prevRune)
	}
	return unpacked, nil
}

func removeLastRune(unpacked *strings.Builder) *strings.Builder {
	runeArr := []rune(unpacked.String())
	runeArr = runeArr[:len(runeArr)-1]
	unpacked.Reset()
	unpacked.WriteString(string(runeArr))
	return unpacked
}

func handleEscapeSlash(r []rune, unpacked *strings.Builder, i int) (*strings.Builder, error) {
	nextRune := r[i+1]
	_, ok := nums[nextRune]
	if !ok && nextRune != '\\' {
		return unpacked, ErrInvalidString
	}
	unpacked.WriteRune(nextRune)
	return unpacked, nil
}
