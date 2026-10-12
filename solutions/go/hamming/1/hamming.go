package hamming

import "errors"


func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length must be the same")
		}

	count := 0

	for i := 0 ; i < len(a); i++ {
		if a[i] != b[i] {
			count += 1
		}
	}
	return count, nil
}
