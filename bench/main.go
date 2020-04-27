// package main

// import "fmt"

// func main() {

// }

// type someinterface interface {
// 	echo()
// }

// type s1 struct {
// 	a int
// }

// func (s s1) echo() {
// 	fmt.Println(s.a)
// }

// func (s s2) echo() {
// 	fmt.Println(s.a)
// }
// func (s s3) echo() {
// 	fmt.Println(s.a)
// }
// func (s s4) echo() {
// 	fmt.Println(s.a)
// }
// func (s s5) echo() {
// 	fmt.Println(s.a)
// }
// func (s s6) echo() {
// 	fmt.Println(s.a)
// }
// func (s s7) echo() {
// 	fmt.Println(s.a)
// }
// func (s s8) echo() {
// 	fmt.Println(s.a)
// }
// func (s s10) echo() {
// 	fmt.Println(s.a)
// }
// func (s s9) echo() {
// 	fmt.Println(s.a)
// }

// type s2 struct {
// 	a int
// }
// type s3 struct {
// 	a int
// }

// type s4 struct {
// 	a int
// }

// type s5 struct {
// 	a int
// }

// type s6 struct {
// 	a int
// }

// type s7 struct {
// 	a int
// }

// type s8 struct {
// 	a int
// }

// type s10 struct {
// 	a int
// }

// type s9 struct {
// 	a int
// }

// func makeInstances() []someinterface {
// 	a1, a2, a3, a4, a5, a6, a7, a8, a9, a10 := s1{1}, s2{1}, s3{1}, s4{1}, s5{1}, s6{1}, s7{1}, s8{1}, s9{1}, s10{1}
// 	test := []someinterface{a1, a2, a3, a4, a5, a6, a7, a8, a9, a10}
// 	for i := 0; i < 10; i++ {
// 		test = append(test, test...)
// 	}
// 	return test
// }

// func executeSwitch([]someinterface) {
// }

package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(23, 12))
}
func uniquePaths(row, col int) int {

	n := uint64(row + col - 2)
	var k1 uint64
	if row < col {
		k1 = uint64(row - 1)
	} else {
		k1 = uint64(col - 1)
	}

	down := uint64(1)
	up := uint64(1)
	for i := uint64(1); i <= k1; i++ {
		down = down * i
	}
	for i := uint64(n - k1 + 1); i <= n; i++ {
		up = up * i
	}
	return int(up / down)
}
