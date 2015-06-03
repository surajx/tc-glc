package main

import (
	"flag"
	"fmt"
	"github.com/surajx/tc-glc/stage1"
	"github.com/surajx/tc-glc/stage2"
)

/*
Stage-1
1:http://community.topcoder.com/tc?module=ProjectDetail&pj=30045937 (Completed)
2:http://community.topcoder.com/tc?module=ProjectDetail&pj=30045938 (Completed)
3:http://community.topcoder.com/tc?module=ProjectDetail&pj=30045939 (Pending)
4:http://community.topcoder.com/tc?module=ProjectDetail&pj=30046010 (Pending)
Stage-2
1:http://community.topcoder.com/tc?module=ProjectDetail&pj=30046011 (Pending)
2:http://community.topcoder.com/tc?module=ProjectDetail&pj=30046225 (Pending)
3:http://community.topcoder.com/tc?module=ProjectDetail&pj=30046224 (Pending)
*/
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
