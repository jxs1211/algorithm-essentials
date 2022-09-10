package main

import "fmt"

func main() {
	grow()
}

func grow() []int {
	arr := [5]int{5, 6, 7, 8, 9}
	slice := arr[2:4]
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 5)
	fmt.Println(len(slice), cap(slice))
	slice = append(slice, 6)
	fmt.Println(len(slice), cap(slice))
	return slice
}
