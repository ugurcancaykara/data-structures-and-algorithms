package main

import "fmt"

func main() {
	exampleList := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println("list before removing more than 2 duplicates")
	fmt.Println(exampleList)

	result := removeDuplicates(&exampleList)
	fmt.Println("unique number of characters(unique means for this question -> have 1 or 2 element of same number): ", result)
	// answer will be the first `result` number of elements of the array, don't care about the rest of it
	fmt.Println("list after removing more than 2 duplicates")
	fmt.Println(exampleList)

}

// solution 1
func removeDuplicates(nums *[]int) int {
	if len(*nums) == 0 {
		return 0
	}

	index := 0

	for i := range *nums {
		if index < 2 || (*nums)[i] != (*nums)[index-2] {
			(*nums)[index] = (*nums)[i]
			index++
		}
	}

	return index
}

// solution 2

//func removeDuplicates(nums []int) int {
//	l, r := 0, 0
//
//	for r < len(nums) {
//		count := 1
//
//		for r+1 < len(nums) && nums[r] == nums[r+1] {
//			count += 1
//			r++
//		}
//
//		for i := 0; i < min(2, count); i++ {
//			nums[l] = nums[r]`
//			l++
//		}
//
//		r++
//	}
//	return l
//}
//
//func min(a, b int) int {
//	if a < b` {
//		return a
//	}
//	return b
//}
