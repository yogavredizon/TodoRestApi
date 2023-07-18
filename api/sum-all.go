package api

import "errors"

func SumAll(a ...int) (int, error) {
	total := 0

	if len(a) > 1 {
		for i := 1; i <= len(a); i++ {
			total += i
		}
		return total, nil
	} else {
		return 0, errors.New("can't fill the argument with 0 value")
	}
}
