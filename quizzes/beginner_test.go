package quizzes

import "testing"

// https://tour.go-zh.org/basics/4  中文
// https://tour.golang.org/basics/4 English
func Test4(t *testing.T) {

	add := func(x int, y int) int {
		return x + y
	}

	// 实现一个减法函数
	// implement function subtraction to pass the test
	// sub := ?

	result := sub(add(1, 2), 3)
	if result != 0 {
		t.FailNow()
	}
}
