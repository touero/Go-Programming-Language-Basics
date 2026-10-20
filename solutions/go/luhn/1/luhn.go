package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	newId := strings.ReplaceAll(id, " ", "")

	if len(newId) <= 1 {
		return false
	}
	
	sum := 0
	index := 0

	for i := len(newId) - 1; i>= 0; i-- {
		temp := rune(newId[i])

		if !unicode.IsDigit(temp) {
			return false
		}

		tempInt := int(temp - '0')

		if index % 2 == 1 {
			tempInt *= 2
			
			if tempInt > 9 {
				tempInt -= 9
			}
		}

		sum += tempInt
		index++
	}

	return sum % 10 == 0
}
