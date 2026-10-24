package phonenumber

import (
	"errors"
	"fmt"
)

func Number(phoneNumber string) (string, error) {
	var b []byte
	for i := 0; i < len(phoneNumber); i++ {
		c := phoneNumber[i]
		if c >= '0' && c <= '9' {
			b = append(b, c)
		}
	}
	digits := string(b)

	if len(digits) == 11 {
		if digits[0] != '1' {
			return "", errors.New("invalid country code")
		}
		digits = digits[1:]
	} else if len(digits) != 10 {
		return "", errors.New("invalid phone number length")
	}

	if digits[0] < '2' || digits[0] > '9' {
		return "", errors.New("invalid area code")
	}
	if digits[3] < '2' || digits[3] > '9' {
		return "", errors.New("invalid exchange code")
	}

	return digits, nil
}

func AreaCode(phoneNumber string) (string, error) {
	result, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return result[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	result, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", result[0:3], result[3:6], result[6:]), nil
}
