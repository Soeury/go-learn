package main

import (
	a "go_learn/learn02/test61/temp" //  这里 a 是 temp 包的别名 ， 可以直接使用

	"fmt"
)

func main() {

	ret := a.Add(1, 2)

	fmt.Println(ret)
}
