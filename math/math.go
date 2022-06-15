package math

func Average(s []float64) float64 {
	avg := float64(0)

	for i := 0; i < len(s); i++ {
		avg += s[i]
	}

	return avg / float64(len(s))
}
