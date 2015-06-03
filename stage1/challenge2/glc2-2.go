package main

import (
	"fmt"
	"sort"
)

func minFloors(M int, heights []int) int {
	sort.Ints(heights)
	minCost := (heights[len(heights)-1] - heights[0]) * M
	for idx, _ := range heights {
		if idx+M <= len(heights) {
			var MBuildings = heights[idx : idx+M]
			var MMax = MBuildings[M-1]
			sliceCost := 0
			for _, val := range MBuildings[:M-1] {
				sliceCost += MMax - val
			}
			if sliceCost < minCost {
				minCost = sliceCost
			}
		} else {
			break
		}
	}
	return minCost
}

type problemData struct {
	maxJump         int
	buildingHeights []int
}

var problemDataSet []problemData

func init() {
	problemDataSet = []problemData{
		{2, []int{1, 2, 1, 4, 3}},
		{3, []int{1, 3, 5, 2, 1}},
		{1, []int{43, 19, 15}},
		{3, []int{19, 23, 9, 12}},
		{12, []int{25, 18, 38, 1, 42, 41, 14, 16, 19, 46, 42, 39,
			38, 31, 43, 37, 26, 41, 33, 37, 45, 27, 19, 24, 33, 11, 22, 20, 36, 4, 4}},
	}
}

func main() {
	for idx, data := range problemDataSet {
		fmt.Printf("Problem No #%d, Minimium Floors: %d\n", idx,
			minFloors(data.maxJump, data.buildingHeights))
	}
}
