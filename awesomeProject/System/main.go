//  这段代码演示了如何使用堆接口构建一个整数堆。
package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("./lib/test")
	if err != nil {
		fmt.Println("error", err)
	}

}
