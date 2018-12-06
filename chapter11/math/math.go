package math

// Avarage - Finds the avarage of a series of numbers
func Avarage(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return 0 + (total / float64(len(xs)))
}
