package datastructuretest

import (
	"fmt"
	"testing"
)

func TestCopySlice(t *testing.T) {
	path := []int{1, 2, 3}
	temp := make([]int, 1)
	copy(temp, path)
	fmt.Print(path, temp)
}

func TestDivide(t *testing.T) {
	a := 10
	b := 3
	fmt.Println(a / b)
}
