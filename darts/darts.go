package darts

// Score returns the points recieved by a dart position
func Score(x float64, y float64) int {
	// Get the cartesian product by x^2 + y^2
	input := x*x + y*y

	// The product must be lower or equal to the radius^2 for each circle
	if input <= 1*1 {
		return 10
	}
	if input <= 5*5 {
		return 5
	}
	if input <= 10*10 {
		return 1
	}
	return 0
}
