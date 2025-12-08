package hw02unpackstring

import (
	"errors"
	"strings"

	"github.com/rivo/uniseg"
)

var (
	ErrInvalidString = errors.New("invalid string")
	nums             = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
)

func Unpack(s string) (string, error) {
	unpacked := &strings.Builder{}
	gr := uniseg.NewGraphemes(s)
	glifs := convertToGlifs(gr)
	unpacked, err := unpackGlifs(glifs, unpacked)
	if err != nil {
		return unpacked.String(), err
	}

	return unpacked.String(), nil
}

func convertToGlifs(gr *uniseg.Graphemes) []string {
	var glifs []string
	for gr.Next() {
		glifs = append(glifs, gr.Str())
	}
	return glifs
}

func unpackGlifs(glifs []string, unpacked *strings.Builder) (*strings.Builder, error) {
	for i := 0; i < len(glifs); i++ {
		num, ok := nums[glifs[i]]
		switch {
		case ok && i > 0:
			unpacked, err := handleNum(glifs, unpacked, num, i)
			if err != nil {
				return unpacked, err
			}

		case func() bool {
			curGlif := glifs[i]
			return curGlif == "\\"
		}():
			unpacked, err := handleEscapeSlash(glifs, unpacked, i)
			if err != nil {
				return unpacked, err
			}
			i++

		case ok && i == 0:
			return unpacked, ErrInvalidString

		default:
			unpacked.WriteString(glifs[i])
		}
	}
	return unpacked, nil
}

func handleNum(glifs []string, unpacked *strings.Builder, num int, i int) (*strings.Builder, error) {
	prevRune := glifs[i-1]
	isLastRune := i == len(glifs)-1
	if !isLastRune {
		nextGlif := glifs[i+1]
		_, ok := nums[nextGlif]
		if ok {
			return unpacked, ErrInvalidString
		}
	}
	if num == 0 {
		unpacked = removeLastRune(unpacked, i)
	}
	for i := 0; i < num-1; i++ {
		unpacked.WriteString(prevRune)
	}
	return unpacked, nil
}

func removeLastRune(unpacked *strings.Builder, i int) *strings.Builder {
	gr := uniseg.NewGraphemes(unpacked.String())
	glifs := convertToGlifs(gr)
	unpacked.Reset()
	if len(glifs) == 1 {
		return unpacked
	}
	glifs = glifs[:len(glifs)-1]
	for _, g := range glifs {
		unpacked.WriteString(g)
	}
	return unpacked
}

func handleEscapeSlash(glifs []string, unpacked *strings.Builder, i int) (*strings.Builder, error) {
	nextGlif := glifs[i+1]
	_, ok := nums[nextGlif]
	if !ok && nextGlif != "\\" {
		return unpacked, ErrInvalidString
	}
	unpacked.WriteString(nextGlif)
	return unpacked, nil
}
