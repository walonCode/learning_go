package mathutils


//take an array of float64 and return the maximum number in that array
func Max(xs []float64)float64{
	max := xs[0]
	for _, x := range xs[1:] {
		if x > max {
			max = x
		}
	}
	return max
}

//take an array of float64 and return the minimum number in that array
func Min(xs []float64)float64{
	min := xs[0]
	for _, x := range xs[1:]{
		if x < min{
			min = x
		}
	}
	return min
}


//takes in an array and return the sum of that array
func Sum(xs []float64)float64{
	total := 0.0
	for _, x := range xs{
		total += x
	}
	return total
}