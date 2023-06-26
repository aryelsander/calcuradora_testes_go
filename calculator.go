package calculator

func Sum(values ...float64) float64 {
	total_sum := 0.0
	for _, v := range values {
		total_sum += v
	}

	return total_sum
}

func Subtract(values ...float64) float64 {
	total_subtract := values[0]
	for i := 1; i < len(values); i++ {
		total_subtract -= values[i]
	}

	return total_subtract
}

func Multiply(values ...float64) float64 {
	total_multiply := 1.0
	for _, v := range values {
		total_multiply *= v
	}
	return total_multiply
}

func Divide(values ...float64) float64 {
	total_division := values[0]

	for i := 1; i < len(values); i++ {
		total_division /= values[i]
	}

	return total_division
}
