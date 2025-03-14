package math

//take two input and add's them together
func Add(a, b int) int {
	return a + b	
}


//take to input and mulitply them together
func Average(xs []float64) float64 {
	total := 0.0
	for _, num := range xs {
		total += num
	}
	return total / float64(len(xs))
}