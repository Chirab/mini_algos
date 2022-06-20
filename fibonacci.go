package main

import "fmt"

func removeFirstElement(tab []int) []int {
	newSlice := make([]int, 0, 1)
        return  append(newSlice, tab[1:][0])
}

func fibonacci(s int) []int {
	fiboSlice := make([]int, 0, 1)
	// Temporary slice where two values will be stored in order to sum them
        tmpSlice := make([]int, 0)
        for i:= 0; i < s; i++ {
		if len(tmpSlice) < 2 {
			tmpSlice = append(tmpSlice, i)
		}
		
                if len(tmpSlice) == 2 {
                        sum := tmpSlice[0] + tmpSlice[1]
                        tmpSlice = append(removeFirstElement(tmpSlice), sum)
			fiboSlice = append(fiboSlice, tmpSlice[0])
                }
        }
	
	return fiboSlice
}

func main() {
	result := fibonacci(1000)
	fmt.Println(result)
}
