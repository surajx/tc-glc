package challenge3

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/surajx/tc-glc/stage1/challenge3/dice"
)

type problemData struct {
	rolls []string
}

func runMinimumFaces(rw http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	problemDataString := req.FormValue("r")
	problemDataSlice := []string{}
	for _, roll := range strings.Split(problemDataString, ",") {
		problemDataSlice = append(problemDataSlice, roll)
	}
	io.WriteString(rw, strconv.Itoa(dice.MinimumFaces(problemDataSlice)))
}

func Challenge3() {

	problemDataSet := []problemData{
		{[]string{"137", "364", "115", "724"}},
		{[]string{"1112", "1111", "1211", "1111"}},
		{[]string{"24412", "56316", "66666", "45625"}},
		{[]string{"931", "821", "156", "512", "129", "358", "555"}},
		{[]string{"3", "7", "4", "2", "4"}},
		{[]string{"281868247265686571829977999522", "611464285871136563343229916655",
			"716739845311113736768779647392", "779122814312329463718383927626",
			"571573431548647653632439431183", "547362375338962625957869719518",
			"539263489892486347713288936885", "417131347396232733384379841536"}},
	}

	for idx, data := range problemDataSet {
		fmt.Printf("Input Data: #%d, Minimium Faces: %d\n", idx, dice.MinimumFaces(data.rolls))
	}

	http.HandleFunc("/stage1/c3", runMinimumFaces)
}
