package main

import (
	"errors"
	"fmt"
)

func getTheBestEngineer(names []string, votes []string) string {
	votesCount := make([]int, len(names))
	getProfIdxFunc := func(name string) (int, error) {
		for idx, profName := range names {
			if name == profName {
				return idx, nil
			}
		}
		return -1, errors.New("Invalid Professor name")
	}
	for idx, profName := range names {
		if profName == votes[idx] {
			continue
		}
		voteIdx, err := getProfIdxFunc(votes[idx])
		if err != nil {
			return err.Error()
		}
		votesCount[voteIdx] += 1
	}
	max := votesCount[0]
	maxIdx := 0
	for idx, value := range votesCount {
		if value > max {
			max = value
			maxIdx = idx
		}
	}
	if max == 0 {
		return ""
	}
	for idx, vc := range votesCount {
		if vc == max && idx != maxIdx {
			return ""
		}
	}
	return names[maxIdx]
}

type problemDataSet struct {
	profNames []string
	profVotes []string
}

func main() {
	problemData := []problemDataSet{
		{[]string{"Toshi", "Jun", "Teru", "Chihiro"}, []string{"Jun", "Chihiro", "Toshi", "Toshi"}},
		{[]string{"Toshi", "Jun", "Teru", "Chihiro"}, []string{"Teru", "Teru", "Jun", "Jun"}},
		{[]string{"Toshi", "Jun", "Teru", "Chihiro"}, []string{"Toshi", "Toshi", "Jun", "Jun"}},
		{[]string{"Toshi", "Jun"}, []string{"Toshi", "Jun"}},
		{[]string{"PhpLove", "phplove", "phpLove", "Phplove"}, []string{"phpLove", "phpLove", "phpLove", "PhpLove"}},
	}

	for idx, data := range problemData {
		fmt.Printf("Problem No %d, Winning Prof: %s\n", idx,
			getTheBestEngineer(data.profNames, data.profVotes))
	}
}
