package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

const escapeCharacter = '\\'

const (
	stringStart = iota
	afterNumber
	afterEscapeCharacter
	afterRegularCharacter
)

const (
	number = iota
	regular
	escape
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	var prevChar string
	position := stringStart

	stringBuilder := strings.Builder{}

	for _, char := range packedString {
		if position == afterEscapeCharacter {
			if unicode.IsNumber(char) || char == escapeCharacter {
				position = afterRegularCharacter
				prevChar = string(char)
				continue
			}

			return "", ErrInvalidString
		}

		switch getCharType(char) {
		case number:
			if position == stringStart || position == afterNumber {
				return "", ErrInvalidString
			}

			position = afterNumber

			repeatNumber, _ := strconv.Atoi(string(char))
			stringBuilder.WriteString(strings.Repeat(prevChar, repeatNumber))
			prevChar = ""
		case escape:
			position = afterEscapeCharacter
			stringBuilder.WriteString(prevChar)
			prevChar = ""
		case regular:
			stringBuilder.WriteString(prevChar)
			prevChar = string(char)
			position = afterRegularCharacter
		}
	}

	if position == afterEscapeCharacter {
		return "", ErrInvalidString
	}

	stringBuilder.WriteString(prevChar)

	return stringBuilder.String(), nil
}

func getCharType(char rune) int {
	if unicode.IsNumber(char) {
		return number
	}

	switch char {
	case escapeCharacter:
		return escape
	default:
		return regular
	}
}
