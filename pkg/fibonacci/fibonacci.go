package fibonacci

// Sequence function calculates the value of the fibonacci sequence.
// It returns the value of a function from the number n.
func Sequence(n int) int {
	if n == 0 { return 0 }
	if n == 1 { return 1 }

	if n < 0 {
		return Sequence(n+2) - Sequence(n+1)
	} else {
		return Sequence(n-1) + Sequence(n-2)
	}
}
