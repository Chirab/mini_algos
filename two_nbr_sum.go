package main

import "fmt"

// Find two numbers in a list that sum to a given amount
func findTwoNumberSum(data []int, amount int) []int {
	res := make([]int, 2)
	for _, v := range data {
		tmp := v
		for _, j := range data {
			sum := tmp + j
			if sum == amount {
				res[0] = tmp
				res[1] = j
			}
		}
	}

	return res
}

func main() {
	data := []int{8, 3, 4, 19, 39, 21, 33}
	amount := 7
	result := findTwoNumberSum(data, amount)
	fmt.Println(result)
}
