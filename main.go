package main

import (
	"fmt"
	"kata/files"
)

func main() {
	fmt.Println(files.MergeArrays([]int{1, 2, 4345, 3}, []int{1, 2, 3, 4, 5}))
}
