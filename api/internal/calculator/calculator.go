package calculator

// Calculator contains a slice of numbers to perform basic operations with
type Calculator struct {
	Numbers []int `json:"numbers"`
}

// Total is a function that will take in a slice of integers,
// and perform addition on all the values to find their total value.
func (c *Calculator) Total() int {
	var (
		sum int
	)
	for _, v := range c.Numbers {
		sum += v
	}
	return sum
}
