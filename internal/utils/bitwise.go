package utils

import (
	"fmt"
	"strconv"
)

func ToBinary(num int) string {
	rs := ""
	for num > 0 {
		rs = fmt.Sprintf("%d%s", num%2, rs)
		num = num / 2
	}
	return rs
}

func AddBinary(a string, b string) string {
	if len(a) < len(b) {
		b, a = a, b
	}
	sizeA := len(a) - 1
	sizeB := len(b) - 1
	var remember uint8 = 0
	rs := ""
	for i := 0; i <= sizeA; i++ {
		bitA := a[sizeA-i] - 48
		var bitB uint8 = 0
		if sizeB-i >= 0 {
			bitB = b[sizeB-i] - 48
		}
		bit := bitA ^ bitB ^ remember
		remember = (bitA & bitB) | (bitA & remember) | (bitB & remember)
		rs = strconv.Itoa(int(bit)) + rs
	}
	if remember == 1 {
		rs = "1" + rs
	}
	return rs
}
