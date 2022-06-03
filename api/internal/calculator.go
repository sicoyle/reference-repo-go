package internal

// Total is a function that will take in a slice of integers,
// and perform addition on all the values to find their total value.
func (d *Digits) Total() int {
	var (
		sum int
	)
	for _, v := range d.Numbers {
		sum += v
	}
	return sum
}
