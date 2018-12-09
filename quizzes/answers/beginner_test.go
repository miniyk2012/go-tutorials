package quizzes

import "testing"

// https://tour.go-zh.org/basics/4  中文
// https://tour.golang.org/basics/4 English
func Test4(t *testing.T) {

	add := func(x int, y int) int {
		return x + y
	}

	sub := func(x int, y int) int {
		return x - y
	}

	result := sub(add(1, 2), 3)
	if result != 0 {
		t.FailNow()
	}
}

// https://go-tour-zh-tw.appspot.com/basics/6 中文
// https://tour.golang.org/basics/6           English
func Test6(t *testing.T) {

	// Implement this function such that it returns the length of each parameter in the same order.
	strLen := func(str1, str2, str3 string) (int, int, int) {
		return len(str1), len(str2), len(str3)
	}

	l1, l2, l3 := strLen("a", "ab", "abc")
	if l1 != 1 || l2 != 2 || l3 != 3 {
		t.FailNow()
	}
}

// http://go-tour-zh.appspot.com/basics/13 中文
// https://tour.golang.org/basics/13       English
func Test13(t *testing.T) {

	// Find the quotient, reminder and module of x / y.
	div := func(x, y int) (quotient float64, reminder, modulo int) {
		quotient = float64(x) / float64(y)
		reminder = x % y
		if reminder < 0 {
			modulo = reminder + y
			return
		}
		modulo = reminder
		return
	}

	q, r, m := div(-21, 4)
	if q != -5.25 {
		t.FailNow()
	}
	if r != -1 {
		t.FailNow()
	}
	if m != 3 {
		t.FailNow()
	}

}
