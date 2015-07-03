package stage1

import (
	"fmt"
	"io"
	"net/http"

	"github.com/surajx/tc-glc/stage1/challenge1"
	"github.com/surajx/tc-glc/stage1/challenge2"
	"github.com/surajx/tc-glc/stage1/challenge3"
	"github.com/surajx/tc-glc/stage1/challenge4"
)

const (
	TOTAL_CHALLEGE_COUNT    = 6
	NETWORK_CHALLENGE_COUNT = 3
)

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func Stage1() {
	doneChan := make(chan bool)
	stopChan := make(chan bool)
	go func() {
		for i := 0; i < TOTAL_CHALLEGE_COUNT; i++ {
			<-doneChan
		}
		stopChan <- true
	}()
	fmt.Println("####Executing Challenge 1####")
	challenge1.Challenge1(doneChan)
	fmt.Println("\n####Executing Challenge 2 Problem 1####")
	challenge2.Challenge2P1(doneChan)
	fmt.Println("\n####Executing Challenge 2 Problem 2####")
	challenge2.Challenge2P2(doneChan)
	fmt.Println("\n####Executing Challenge 3####")
	challenge3.Challenge3()
	fmt.Println("\n####Executing Challenge 4####")
	challenge4.Challenge4()
	challenge4.Challenge4_tcp()
	http.HandleFunc("/stop", func(rw http.ResponseWriter, req *http.Request) {
		//Need to use custom code to gracefully stop http server
		//http://www.hydrogen18.com/blog/stop-listening-http-server-go.html
		//http://rcrowley.org/articles/golang-graceful-stop.html
		//stopping all http server based challenges
		for i := 0; i < NETWORK_CHALLENGE_COUNT; i++ {
			doneChan <- true
		}
		io.WriteString(rw, "Server Stopped.")
	})
	go func() {
		err := http.ListenAndServe(":8899", nil)
		chk(err)
	}()
	fmt.Println("Visit http://localhost:8899/stop to stop server and exit application.")
	<-stopChan
}
