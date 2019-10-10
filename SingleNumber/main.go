package main

import "fmt"

func main() {
	arr := []int{12, 7, 7, 78, 78, 4, 4, 6, 6, 13, 12}
	fmt.Println(singleNumber(arr))
	fmt.Println(singleNumber2(arr))
}

func singleNumber(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum = sum ^ nums[i]
	}
	return sum
}

func singleNumber2(nums []int) int {
	sum := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		sum[nums[i]]++
	}
	for key, v := range sum {
		if v == 1 {
			return key
		}
	}
	return 0
}
