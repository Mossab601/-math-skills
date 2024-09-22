package main

import (
	"fmt"
	"function/Function"
)



func main() {
	var nums function.Numbers

	err := nums.ReadFromFile("data.txt")
	if err != nil {
		fmt.Println(" error in read :", err)
		return
	}

	average := nums.Average()
	median:=nums.Median()
	variance:=nums.Variance()
	standiv:=nums.StandardDeviation()
	fmt.Println(" Average:", average)
	fmt.Println(" Median:", median)
	fmt.Println(" Variance:", variance)
	fmt.Println(" StandardDeviation:", standiv)

}
