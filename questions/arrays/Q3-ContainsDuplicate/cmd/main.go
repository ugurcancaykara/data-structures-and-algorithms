package main

import "fmt"

func main() {

	exampleList := []int{1, 2, 3, 1}
	result := containsDuplicate(exampleList)
	fmt.Println(result)

}

func containsDuplicate(nums []int) bool {

	keyMap := make(map[int]int)

	for i := range nums {
		if _, ok := keyMap[nums[i]]; ok {
			// if it exists in the map then it means it already counted
			return true
		}
		keyMap[nums[i]] = 1
	}
	fmt.Println(keyMap)
	return false
}
