package Template

import "fmt"

//以下都是供本包内的函数进行调用的
func Template(a, b int) int {
	fmt.Println("Template")
	return a / b
}
