package pkg

import (
	"C"
	"fmt"
	"math/big"
)

//export Ping
func Ping() {
	fmt.Println("Pong")
}

//export Factorial
func Factorial(num int) {
	var fact big.Int
	fact.MulRange(1, int64(num))
	fmt.Println(fact.String())
}

//export TotalAdd
func TotalAdd(start, end int) int {
	var total int = 0
	for i := start; i <= end; i++ {
		total += i
	}
	return total
}
