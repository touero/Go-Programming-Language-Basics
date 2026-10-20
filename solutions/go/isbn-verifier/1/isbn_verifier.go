package isbnverifier

import "unicode"

func IsValidISBN(isbn string) bool {
	result := 0
	count := 0

	for i := 0; i < len(isbn); i++ {
		temp := rune(isbn[i])

		if temp == '-' || temp == ' ' {
			continue
		}

		count++

		if temp == 'X' || temp == 'x' {
			if count != 10 {
				return false
			}
			result += 10 * (10 - (count - 1))
			continue
		}

		if !unicode.IsDigit(temp) {
			return false
		}

		result += int(temp-'0') * (10 - (count - 1))
	}

	if count != 10 {
		return false
	}

	if result%11 == 0 {
		return true
	}

	return false
}
