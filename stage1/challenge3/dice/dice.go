package dice

import (
	"fmt"
	"sort"
	"strconv"
)

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//MinimumFaces returns the minimum number of faces
//that can be expected from the the given set of rolls.
func MinimumFaces(rolls []string) int {
	rollsInt := [][]int{}
	for _, roll := range rolls {
		rollInt := []int{}
		for _, pip := range roll {
			faceValue, err := strconv.Atoi(string(pip))
			chk(err)
			rollInt = append(rollInt, faceValue)
		}
		sort.Sort(sort.Reverse(sort.IntSlice(rollInt)))
		rollsInt = append(rollsInt, rollInt)
	}
	minFacesRoll := make([]int, len(rollsInt[0]))
	for _, roll := range rollsInt {
		for idx, faceVal := range roll {
			if faceVal > minFacesRoll[idx] {
				minFacesRoll[idx] = faceVal
			}
		}
	}
	sum := 0
	for _, faceVal := range minFacesRoll {
		sum += faceVal
	}
	return sum
}
