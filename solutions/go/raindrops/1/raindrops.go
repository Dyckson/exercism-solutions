package raindrops

import strconv "strconv"


func Convert(number int) string {
	isValid := false
	result := ""
	if number%3 == 0 {
		isValid = true
		result += "Pling"
	}

	if number%5 == 0 {
		isValid = true
		result += "Plang"
	}

	if number%7 == 0 {
		isValid = true
		result += "Plong"
	}
	if !isValid {
		return strconv.Itoa(number)
	}
	return result
}

