package examples

import "fmt"

func Pow(muns []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, m := range muns {
			out <- m * m
		}
		close(out)
	}()

	return out
}

func RunPipeLineExample() {
	fmt.Println(Pow([]int{1, 2, 3, 4, 5}))
}
