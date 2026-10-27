package hamming
import "fmt"


func Distance(a, b string) (int, error) {
	cauntDistance := 0
    if len(a) != len(b) {
        return 0, fmt.Errorf("strings must have the same length")
    }
	for i, l := range a {
		if l != rune(b[i]) {
			cauntDistance++
		}
	}

	return cauntDistance, nil
}
