package hw02unpackstring

import (
	"errors"
	"strings"

	"github.com/rivo/uniseg"
)

var (
	ErrInvalidString = errors.New("invalid string")
)

func Unpack(s string) (string, error) {
	unpacked := &strings.Builder{}
	gr := uniseg.NewGraphemes(s)
	glifs := make([]string, 0, len(s))
	for gr.Next() {
		glifs = append(glifs, gr.Str())
	}
	unpacked, err := unpackGlifs(glifs, unpacked)
	if err != nil {
		return unpacked.String(), err
	}

	return unpacked.String(), nil
}

func unpackGlifs(glifs []string, unpacked *strings.Builder) (*strings.Builder, error) {
	for i := 0; i < len(glifs); i++ {
		num := glifs[i]
		isNum := len(num) == 1 && num[0] >= '0' && num[0] <= '9'
		switch {
		case isNum && i > 0:
			digit := int(num[0] - '0')
			unpacked, err := handleNum(glifs, unpacked, digit, i)
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

		case isNum && i == 0:
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
		isNum := len(nextGlif) == 1 && nextGlif[0] >= '0' && nextGlif[0] <= '9'
		if isNum {
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
	gr := uniseg.NewGraphemes(unpacked.String())
	glifs := make([]string, 0, len(unpacked.String()))
	for gr.Next() {
		glifs = append(glifs, gr.Str())
	}
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
	isNum := len(nextGlif) == 1 && nextGlif[0] >= '0' && nextGlif[0] <= '9'
	if !isNum && nextGlif != "\\" {
		return unpacked, ErrInvalidString
	}
	unpacked.WriteString(nextGlif)
	return unpacked, nil
}
