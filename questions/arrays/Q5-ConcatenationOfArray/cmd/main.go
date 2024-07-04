package main

import "fmt"

func main() {
	exampleList := []int{1, 3, 2, 1}
	fmt.Println("list before concatenation of array")
	fmt.Println(exampleList)

	result := getConcatenation(exampleList)
	fmt.Println(result)
	fmt.Println("list after concatenation of array")
	fmt.Println(exampleList)

}

func getConcatenation(nums []int) []int {
	// create ans array which is length 2n
	ans := make([]int, len(nums)*2)

	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		ans[i+len(nums)] = nums[i]
	}
	return ans
}
