package unescapeString

import (
	"errors"
	"strconv"
	"unicode"
)

// UnescapeString выполняет распаковку строки с учетом escape-последовательностей
func UnescapeString(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	var result []rune
	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length; {
		switch {
		case runes[i] == '\\':
			if i+1 >= length {
				return "", errors.New("escape character at end of string")
			}
			result = append(result, runes[i+1])
			i += 2

		case unicode.IsDigit(runes[i]):
			if len(result) == 0 {
				return "", errors.New("digit without preceding character")
			}

			start := i
			for i < length && unicode.IsDigit(runes[i]) {
				i++
			}
			count, _ := strconv.Atoi(string(runes[start:i]))

			if count <= 0 {
				return "", errors.New("invalid repeat count")
			}

			lastChar := result[len(result)-1]
			if count > 1 {
				result = append(result, make([]rune, count-1)...)
				for j := 1; j < count; j++ {
					result[len(result)-j] = lastChar
				}
			}

		default:
			result = append(result, runes[i])
			i++
		}
	}

	if len(result) == 0 {
		return "", errors.New("resulting string is empty")
	}

	return string(result), nil
}
