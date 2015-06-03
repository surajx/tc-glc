package challenge1

import (
	"fmt"
)

func minimumCost(altitudeSlice []int) int {
	prevAltitude := -1
	cost := 0
	for _, altitude := range altitudeSlice {
		if prevAltitude == -1 {
			prevAltitude = altitude
			continue
		}
		if prevAltitude < altitude {
			cost += altitude - prevAltitude
		} else {
			prevAltitude = altitude
		}
	}
	return cost
}

func Challenge1() {

	inputData := [][]int{
		[]int{30, 20, 20, 10},
		[]int{5, 7, 3},
		[]int{6, 8, 5, 4, 7, 4, 2, 3, 1},
		[]int{749, 560, 921, 166, 757, 818, 228, 584, 366, 88},
		[]int{712, 745, 230, 200, 648, 440, 115,
			913, 627, 621, 186, 222, 741, 954, 581, 193, 266, 320, 798, 745},
	}

	for _, data := range inputData {
		fmt.Printf("Input Altitude Data: %v, Minimum Cost: %d\n", data, minimumCost(data))
	}
}
