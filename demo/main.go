package main

import (
	"fmt"
	"strconv"
)

func main() {
	intToHexForEVM()
	strToHexForEVM()
}

func intToHexForEVM() {
	var num int64

	//b4e97a0  0xb4e97a0
	num = 189700000

	hexNum := strconv.FormatInt(num, 16)

	fmt.Println("hexadecimal num: ", hexNum)
}

func strToHexForEVM() {
	r := "189700000"
	e, err := strconv.ParseInt(r,10, 64)
	if err != nil {
		fmt.Printf("Parse Error %v", err)
	}

	fmt.Println("hexadecimal num: ", strconv.FormatInt(e, 16))
}
