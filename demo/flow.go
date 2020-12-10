package demo

import (
	"errors"
	"fmt"
)

// Flow : used for demo
func Flow() {
	result, error := div(6, 2)
	fmt.Println(result, error) // 3 <nil>

	result, error = div(6, 0)
	// fmt.Println(result, error)
	if error != nil {
		fmt.Println("error: ", error)
	}
	fmt.Println("success: ", result)

	demoDefer()
}

func div(a, b int) (r int, e error) {
	if b == 0 {
		e = errors.New("can not divide 0")
		return
	}
	r = a / b
	return
}

func demoDefer() {

	fmt.Println("a")
	defer func() {
		fmt.Println("defer 1")
		if error := recover(); error != nil { // 返回panic信息，恢复当前函数或当前函数调用函数中的panic
			fmt.Println("出现了panic，panic的信息为:", error)
		}
	}()
	fmt.Println("b")
	defer func() {
		fmt.Println("defer 2")
	}()
	panic("stop here")
	fmt.Println("will not be reached as panic")

}
