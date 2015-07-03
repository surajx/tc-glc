package challenge4

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strconv"
	"strings"

	"github.com/surajx/tc-glc/stage1/challenge4/survey"
)

func runCountBigShoppersTCP(con net.Conn) {
	inputDataRaw, err := ioutil.ReadAll(con)
	inputDataRawSlice := strings.Split(string(inputDataRaw), "#")
	if len(inputDataRawSlice) != 2 {
		fmt.Println("Invalid data")
		return
	}
	total, err := strconv.Atoi(inputDataRawSlice[0])
	if err != nil {
		fmt.Println(err)
		return
	}
	itemsBought := []int{}
	for _, itemStr := range strings.Split(inputDataRawSlice[1], ",") {
		item, err := strconv.Atoi(itemStr)
		if err != nil {
			fmt.Println(err)
			return
		}
		itemsBought = append(itemsBought, item)
	}
	io.WriteString(con, strconv.Itoa(survey.CountBigShoppers(total, itemsBought))+"\n")
	con.Close()
}

func Challenge4_tcp() {
	tcpListener, err := net.Listen("tcp", ":8099")
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for {
			conn, err := tcpListener.Accept()
			if err != nil {
				tcpListener.Close()
				fmt.Println(err)
				return
			}
			go runCountBigShoppersTCP(conn)
		}
	}()
}
