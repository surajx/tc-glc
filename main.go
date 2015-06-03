package main

import (
	"flag"
	"fmt"
	"github.com/surajx/tc-glc/stage1"
	"github.com/surajx/tc-glc/stage2"
)

func main() {
	var stage int
	flag.IntVar(&stage, "stage", -1, "Stage to execute")
	flag.Parse()
	switch stage {
	case 1:
		stage1.Stage1()
	case 2:
		stage2.Stage2()
	default:
		fmt.Println("Invalida stage")
	}
}
