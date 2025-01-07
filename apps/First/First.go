package First

import "fmt"

//以下都是供本包内的函数进行调用的
func First(a, b int) int {
	fmt.Println("First")
	return a * b
}
