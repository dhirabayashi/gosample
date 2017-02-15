package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/dhirabayashi/gosample"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	checkError(err)

	y, err := strconv.Atoi(os.Args[2])
	checkError(err)

	fmt.Printf("x + y = %d\n", gosample.Add(x, y))
	fmt.Printf("x - y = %d\n", gosample.Subtract(x, y))
	fmt.Printf("x * y = %d\n", gosample.Multiply(x, y))

	answer, rest, err := gosample.Divide(x, y)

	if err != nil {
		fmt.Printf("error occured during muliply: %v\n", err)
	}
	fmt.Printf("x / y: answer = %d, rest = %d\n", answer, rest)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}