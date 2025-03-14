package math

func Add(a, b int) int {
	return a + b	
}

func Average(xs []float64) float64 {
	total := 0.0
	for _, num := range xs {
		total += num
	}
	return total / float64(len(xs))
}