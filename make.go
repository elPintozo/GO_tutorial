package main

import "fmt"


func main() {
	//(tipo, capacidad, longuitud)
	slice := make([]int, 4, 7)
	slice_2 := []int{1,2,3,4}

	copia := make([]int, len(slice_2))

	//se agrega un nuevo elemento 
	slice = append(slice, 3)
	slice = append(slice, 2)

	fmt.Println(slice)

	// es similar a la funcion len para los arreglos
	fmt.Println(cap(slice))//longuitud
	fmt.Println(len(slice))//capacidad

	fmt.Println(copia)//antes

	//solo copia la longuitud de la cual dispone copia
	copy(copia, slice_2)
	fmt.Println(copia)//despues
}