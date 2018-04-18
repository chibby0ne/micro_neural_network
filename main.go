package main

import (
	"fmt"
	"os"

	"github.com/chibby0ne/micro_neural_network/matrix"
)

func main() {
	m, err := matrix.NewMatrix(3, 3)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println(m)
}
