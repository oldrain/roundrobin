package test

import (
	"roundrobin"
	"fmt"
)

func ExampleNRR_Next() {
	rr := new(roundrobin.NRR)
	rr.InitRR(exampleNodes)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", rr.Next().Item)
	}

	// Output: a b c a b c a b c a
}

func ExampleRW_Next() {
	rr := new(roundrobin.RW)
	rr.InitRR(exampleNodes)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", rr.Next().Item)
	}
}

func ExampleWRR_Next() {
	rr := new(roundrobin.WRR)
	rr.InitRR(exampleNodes)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", rr.Next().Item)
	}

	// Output: a a a a a b c a a a
}

func ExampleSW_Next() {
	rr := new(roundrobin.SW)
	rr.InitRR(exampleNodes)

	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", rr.Next().Item)
	}

	// Output: a a b a c a a a a b
}
