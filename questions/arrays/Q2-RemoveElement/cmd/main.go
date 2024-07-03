package main

import "fmt"

func main() {
	exampleList := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println("list before removing element")
	fmt.Println(exampleList)

	result := removeElement(&exampleList, 2)
	fmt.Println("number of characters left: ", result)

	// like in the question, order is not important, just first 5 value of the list is the answer
	fmt.Println("list after removing element")
	fmt.Println(exampleList)
}

func removeElement(nums *[]int, val int) int {

	// define k as a pointer to keep the location of last modified index(modified with a value != val)
	k := 0

	for i := range *nums {
		if (*nums)[i] != val {
			(*nums)[k] = (*nums)[i]
			k += 1
		}
	}

	return k
}
