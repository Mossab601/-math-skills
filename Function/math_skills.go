package function

import (
	"math"
	"sort"
)

// Average
func (n *Numbers) Average() float64 {
	if len(n.Nums) == 0 {
		return 0
	}
	sum := 0
	for _, num := range n.Nums {
		sum += num
	}
	average := float64(sum) / float64(len(n.Nums))
	return average
}

// Median
func (n *Numbers) Median() int {
	if len(n.Nums) == 0 {
		return 0 
	}

	sort.Ints(n.Nums)

	mid := len(n.Nums) / 2

	if len(n.Nums)%2 != 0 {
		return (n.Nums[mid])
	}

	med:= float64(n.Nums[mid-1]+n.Nums[mid]) / 2
	return int(math.Round(med))
}

// Variance
func (n *Numbers) Variance() int {
	if len(n.Nums) == 0 {
		return 0
	}

	average := n.Average()
	sumOfSquares := 0.0

	for _, num := range n.Nums {
		diff := float64(num) - average
		sumOfSquares += float64(diff * diff)
	}
	variance := sumOfSquares / float64(len(n.Nums))
	return int(math.Round(variance))
}

// StandardDeviation
func (n *Numbers) StandardDeviation() int {
	stdiv:= math.Sqrt(float64(n.Variance()))
	return int(math.Round(stdiv))
}