package collatzconjecture

import "errors"

var ErrInvalidInput = errors.New("n deve ser um inteiro maior que zero")

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, ErrInvalidInput
	}

	steps := 0
	for n > 1 {
		steps++

		if isEven(n) {
			n /= 2
			continue
		}

		n = (3 * n) + 1
	}

	return steps, nil
}

func isEven(n int) bool {
	return n%2 == 0
}