package main

import (
	"log"
)

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := containsDuplicate(nums)
	log.Println(nums)
	log.Println(k)

}
func singleNumber(nums []int) int {
	m := map[int]int{}
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	for i, v := range m {
		if v == 1 {
			return i
		}
	}
	return nums[0]
}

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, v := range nums {
		m[v] = m[v] + 1
	}

	for _, v := range m {
		if v > 1 {
			return true
		}
	}

	return false
}
