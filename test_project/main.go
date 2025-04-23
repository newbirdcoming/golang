package main

import (
	"fmt"

	"github.com/newbirdcoming/mathutil"
)

func main() {
	resultAdd := mathutil.Add(2, 4)
	resultMultiply := mathutil.Multiply(2, 4)
	fmt.Printf("2 + 4 = %d\n", resultAdd)
	fmt.Printf("2 * 4 = %d\n", resultMultiply)111
}
