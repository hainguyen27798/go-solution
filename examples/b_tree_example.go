package examples

import (
	"github.com/hainguyen27798/go-solution/internal/utils"
)

func RunBTreeExample() {
	bt := utils.NewBTree(2)
	keys := []int{10, 20, 5, 6, 12, 30, 7}

	for _, k := range keys {
		bt.Insert(k)
	}

	bt.Print()
}
