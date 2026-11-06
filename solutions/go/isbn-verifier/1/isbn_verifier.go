package isbnverifier

import ("strings")


func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0

	for i, r := range isbn {
		peso := 10 - i

		if r == 'X' {
			if i != 9 {
				return false
			}
			sum += 10 * peso
			continue
		}

		if r < '0' || r > '9' {
			return false
		}

		num := int(r - '0')
		sum += num * peso
	}

	return (sum % 11 == 0)
}