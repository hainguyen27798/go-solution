package examples

import (
	"fmt"
	"unsafe"
)

type City uint8

const (
	NewYork City = iota
	London
	Paris
	Mumbai
)

type Person struct {
	currentResidence City
	uniqueID         int64
	passportNumber   int16
}

func RunStructExample() {
	me := Person{
		currentResidence: 1,
		uniqueID:         9248511308,
		passportNumber:   10564,
	}

	fmt.Printf(
		"My Person struct uses %d bytes.\n",
		unsafe.Sizeof(me),
	)
}
