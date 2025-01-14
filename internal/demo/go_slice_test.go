package demo

import (
	"fmt"
	"strings"
	"testing"
)

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s ？
	s = append(s, 100)
	return s
}

func myAppend2(s []int) {
	s[1] = 100
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
}

func myAppend3(s []int) []int {
	fmt.Printf("myAppend 内：append前的切片值：%v\n", s)
	fmt.Printf("myAppend 内：append前的长度：%d，容量：%d\n", len(s), cap(s))
	fmt.Printf("myAppend 内：append前的底层数组指针：%p\n", &s[0])

	s = append(s, 100)

	fmt.Printf("myAppend 内：append后的切片值：%v\n", s)
	fmt.Printf("myAppend 内：append后的长度：%d，容量：%d\n", len(s), cap(s))
	fmt.Printf("myAppend 内：append后的底层数组指针：%p\n", &s[0])
	return s
}

func TestGoSlice(t *testing.T) {
	s := make([]int, 10)
	s[0] = 1
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

func TestGoSlice2(t *testing.T) {
	s := make([]int, 10)
	s[0] = 1
	myAppend2(s)

	fmt.Println(s)
}

func TestGoSlice3(t *testing.T) {
	s := make([]int, 3, 5) // 长度为3，容量为5的切片
	s[0], s[1], s[2] = 1, 2, 3

	fmt.Printf("main 中：原始切片值：%v\n", s)
	fmt.Printf("main 中：原始长度：%d，容量：%d\n", len(s), cap(s))
	fmt.Printf("main 中：原始底层数组指针：%p\n", &s[0])

	newS := myAppend3(s)

	fmt.Printf("\nmain 中：调用后原切片值：%v\n", s)
	fmt.Printf("main 中：调用后长度：%d，容量：%d\n", len(s), cap(s))
	fmt.Printf("main 中：调用后底层数组指针：%p\n", &s[0])
	fmt.Printf("main 中：调用后新切片值：%v\n", newS)
	fmt.Printf("main 中：newS 调用后长度：%d，容量：%d\n", len(newS), cap(newS))
	fmt.Printf("main 中：newS 调用后 原始底层数组指针：%p\n", &newS[0])

	// 让我们看看原切片的底层数组是否真的改变了
	fmt.Printf("\n检查原切片 s 的底层数组（包括未使用的容量部分）：\n")
	fullS := s[:cap(s)] // 创建一个包含所有容量的切片
	for i := 0; i < cap(s); i++ {
		fmt.Printf("s[%d] = %d\n", i, fullS[i])
	}
}

func TestGoSlice4(t *testing.T) {
	var toBoard func(path [][]int) []string
	n := 4
	toBoard = func(path [][]int) []string {
		res := make([]string, 0)
		for i := 0; i < n; i++ {
			temp := make([]string, n)
			q := path[i]
			_, ql := q[0], q[1]
			for j := 0; j < n; j++ {
				if j == ql {
					temp[j] = "Q"
				} else {
					temp[j] = "."
				}
			}
			tempString := strings.Join(temp, "")
			res = append(res, tempString)
		}
		return res
	}
	fmt.Println(toBoard([][]int{{0, 0}, {1, 1}, {3, 3}, {2, 2}}))
}
