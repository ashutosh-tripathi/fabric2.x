package main

import (
	"fmt"
)

func calculate(x int) int {

	return x * 2
}

func main() {

	fmt.Printf("print function executed %d", calculate(3))

}
