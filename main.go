package main

import (
	"fmt"

	"github.com/oanaOM/blockchain-go/ledger"
)

func main() {
	c := ledger.Ledger{}

	for x := range []int{1, 2, 4, 5,7,8,9} {
		hashString := fmt.Sprintf("%v", x)
		c.Add(hashString)

	}

	b, err := c.Get("3")

	fmt.Printf("block [%v]: %v, %v\n", 3, b, err)

}
