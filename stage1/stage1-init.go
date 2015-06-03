package stage1

import (
	"fmt"
	"github.com/surajx/tc-glc/stage1/challenge1"
	"github.com/surajx/tc-glc/stage1/challenge2"
	"github.com/surajx/tc-glc/stage1/challenge3"
)

func Stage1() {
	fmt.Println("####Executing Challenge 1####")
	challenge1.Challenge1()
	fmt.Println("\n####Executing Challenge 2 Problem 1####")
	challenge2.Challenge2P1()
	fmt.Println("\n####Executing Challenge 2 Problem 2####")
	challenge2.Challenge2P2()
	fmt.Println("\n####Executing Challenge 3####")
	challenge3.Challenge3()
}
