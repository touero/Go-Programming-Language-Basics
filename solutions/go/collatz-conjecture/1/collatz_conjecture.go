package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("error input")
	}
	count := 0
	for i := 0; n != 1; i++ {
		if n % 2 == 0 {
			n = n / 2
			count += 1
		} else {
			n = n * 3 + 1
			count += 1
		}
	}
	return count, nil
}
