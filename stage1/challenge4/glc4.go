package challenge4

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/surajx/tc-glc/stage1/challenge4/survey"
)

type problemData struct {
	total       int
	timesBought []int
}

func runCountBigShoppersHTTP(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	totalRaw := req.FormValue("t")
	itemsBoughtRaw := req.FormValue("i")
	total, err := strconv.Atoi(totalRaw)
	if err != nil {
		io.WriteString(rw, "Invalid Total Data")
		return
	}
	itemsBought := []int{}
	for _, itemStr := range strings.Split(itemsBoughtRaw, ",") {
		item, err := strconv.Atoi(itemStr)
		if err != nil {
			io.WriteString(rw, "Invalid Items Bought Data")
			return
		}
		itemsBought = append(itemsBought, item)
	}
	io.WriteString(rw, strconv.Itoa(survey.CountBigShoppers(total, itemsBought)))
}

func Challenge4() {
	problemSet := []problemData{
		{5, []int{3, 3}},
		{100, []int{97}},
		{10, []int{9, 9, 9, 9, 9}},
		{7, []int{1, 2, 3}},
		{5, []int{3, 3, 3}},
		{92, []int{92, 90, 92, 92, 92, 89, 92, 91, 92, 89, 92}},
		{99, []int{99, 99, 99, 99, 99, 96, 99, 99, 99, 99, 98, 99, 99}},
		{64, []int{62, 64, 64, 61, 64, 64, 62, 64, 64, 64, 63, 64}},
	}
	for _, data := range problemSet {
		fmt.Printf("Input Data: %v, Minimium Big Shoppers: %d\n",
			data, survey.CountBigShoppers(data.total, data.timesBought))
	}

	http.HandleFunc("/stage1/c4", runCountBigShoppersHTTP)
}
