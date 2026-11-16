package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
)

var nonDigitsRegexp = regexp.MustCompile(`\D`)

func Number(phoneNumber string) (string, error) {
	cleanNumber := nonDigitsRegexp.ReplaceAllString(phoneNumber, "")

	if len(cleanNumber) == 11 {
		if cleanNumber[0] != '1' {
			return "", errors.New("invalid when 11 digits does not start with a 1")
		}
		cleanNumber = cleanNumber[1:]
	}

	if len(cleanNumber) != 10 {
		return "", errors.New("invalid number length")
	}
    
	if cleanNumber[0] < '2' || cleanNumber[0] > '9' {
		return "", errors.New("area code cannot start with zero or one")
	}

	if cleanNumber[3] < '2' || cleanNumber[3] > '9' {
		return "", errors.New("exchange code cannot start with zero or one")
	}

	return cleanNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return number[:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err 
	}
	
	formatPhoneNumber := fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:])
	return formatPhoneNumber, nil
}
