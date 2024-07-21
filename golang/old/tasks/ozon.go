package main

import (
	"log"
)

func parse(source [][]int) [][]int {
	res := [][]int{}
	
	if len(source) == 0 {
		return nil
	}
	
	start := source[0][0]
	end := source[0][1]
	for i := 1; i < len(source); i++ {
		if end >= source[i][0] {
			if source[i][1] > end {
				end = source[i][1]
			}
		} else {
			res = append(res, []int{start, end})
			start = source[i][0]
			end = source[i][1]
		}
	}
	res = append(res, []int{start, end})

	return res
}

func main() {
	log.Printf("Start")
	
	// [[1, 3], [2, 6], [8, 10]] -> [[1, 6], [8, 10]]
	res := parse([][]int{{1,3}, {2,6}, {3, 7}, {8, 10}, {9, 20}})
	
	log.Printf("res: %v", res)
}
