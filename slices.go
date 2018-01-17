package main

import "fmt"

func main() {

	//Al ser declarados de esta manera, los lugares adquieren valor null(nil)
	var slice []int

	slice_2 := []int{1,2,3,4}

	fmt.Println(slice)
	fmt.Println(slice_2)
	fmt.Println(slice_2[2])
	fmt.Println(len(slice_2))

	//arreglo to slice
	arreglo := [9]int{1,2,3,3,4,5,6,7,8}
	//slicing
	slice_3 := arreglo[:5]
	slice_4 := arreglo[5:7]
	slice_5 := arreglo[4:]
	slice_6 := arreglo[:4]
	fmt.Println(arreglo)
	fmt.Println(slice_3)
	fmt.Println(slice_4)
	fmt.Println(slice_5)
	fmt.Println(slice_6)
}