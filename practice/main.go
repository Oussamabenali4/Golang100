package main

import "fmt"

func main() {
	myMap := make(map[int]int)
	nums := [7]int{2, 2, 1, 1, 1, 2, 2}

	for i := 0; i < len(nums); i++ {
		num := nums[i]
		myMap[num] += 1
	}

	for key, value := range myMap {
		if len(nums)%2 == 0 {
			if value >= len(nums)/2 {
				fmt.Println(key)
			}
		} else {
			if value >= len(nums)/2+1 {
				fmt.Println(key)
			}
		}

	}
}
