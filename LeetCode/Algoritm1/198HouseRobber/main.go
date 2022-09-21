package main

import "fmt"

func main() {
	nums := [5]int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
func rob(nums [5]int) int {
	rob1, rob2 := 0, 0

	for _, n := range nums {
		temp := max(n+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}

	return rob2
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
