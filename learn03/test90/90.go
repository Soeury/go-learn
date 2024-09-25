package main

import "fmt"

// 1. 接口断言
func PrintString(arr interface{}) {
	for _, v := range arr.([]string) {
		fmt.Println(v)
	}
}

func PrintInt(arr interface{}) {
	for _, v := range arr.([]int) {
		fmt.Println(v)
	}
}

// 2. 泛型
func PrintTar[T string | int](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func PrintAny[T any](arr []T) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func main() {

	// 泛型初识

	s1 := make([]string, 2)
	s1 = append(s1, "apple", "grape")

	s2 := []int{25, 36}
	PrintString(s1)
	PrintInt(s2)

	fmt.Println("------")

	PrintTar(s1)
	PrintTar(s2)

	fmt.Println("------")

	PrintAny(s1)
	PrintAny(s2)

}
