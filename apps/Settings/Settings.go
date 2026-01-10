package Settings

import (
	"fmt"
)

// 以下都是供本包内的函数进行调用的
func settings(a, b int) int {
	fmt.Println("settings")
	return a / b
}
