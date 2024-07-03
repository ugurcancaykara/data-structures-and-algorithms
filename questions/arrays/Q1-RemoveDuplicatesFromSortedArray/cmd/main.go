package main

import "fmt"

func main() {
	exampleList := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println("list before removing duplicates")
	fmt.Println(exampleList)

	result := removeDuplicates(&exampleList)
	fmt.Println("unique number of characters: ", result)
	fmt.Println("list after removing duplicates")
	fmt.Println(exampleList)

}

func removeDuplicates(nums *[]int) int {
	// we will have two pointer -> left and right,
	// left will follow unique values
	// right will search in the array

	l := 1
	// we will not start with 0 since first value will be unique already

	for r := range *nums {
		// if it's first element, just continue
		if r == 0 {
			continue
		}

		// otherwise check whether the values
		if (*nums)[r] != (*nums)[r-1] {
			// basically if we found a new unique value
			// then assign it to where we left the left pointer and increase left pointer by one for the new unique value
			(*nums)[l] = (*nums)[r]
			l += 1
		}

	}

	// return unique value count
	return l
}
