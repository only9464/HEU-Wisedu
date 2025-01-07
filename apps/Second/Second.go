package Second

import "fmt"

//以下都是供本包内的函数进行调用的
func Second(a, b int) int {
	fmt.Println("Second")
	return a + b
}
