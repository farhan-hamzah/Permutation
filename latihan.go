package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(path []int, used []bool)

	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			// Salin slice agar tidak berubah
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			// Backtrack
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	used := make([]bool, len(nums))
	backtrack([]int{}, used)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	fmt.Println("Hasil permutasi:")
	for _, p := range result {
		fmt.Println(p)
	}
}
