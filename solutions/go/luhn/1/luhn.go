package luhn

import "strings"


func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
	return (!(len(id) < 2) && IsLuhnValid(id))
}

func IsLuhnValid(cardNumber string) bool {
	sum := 0
	alternate := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if digit < 0 || digit > 9 {
			return false
		}

		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}