package math

// Avarage - Finds the avarage of a series of numbers
func Avarage(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
