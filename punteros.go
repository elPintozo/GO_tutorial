package main

import "fmt"

func main() {
	/*
		- Un puntero es una direccion de memoria
		- en lugar de obtener el valor de la variable, se obtiene su ubicacion de memoria
		- se puede modificar el punto de memoria en donde las variables tienen almacenado su
		valor, logrando asi cambiar su valor sin realizar una declaracion
	*/
	var puntero_1, puntero_2 *int
	entero_1 := 5
	entero_2 := 10

	// & me da acceso a la direccion de memoria en donde se almacena el valor de la variable
	puntero_1 = &entero_1
	puntero_2 = &entero_2

	fmt.Println(puntero_1)
	fmt.Println(puntero_2)

	// * para indicar que quiero el valor alojado dentro de ese esa direccion de memoria
	fmt.Println(*puntero_1)
	fmt.Println(*puntero_2)

	fmt.Println(entero_1)
	fmt.Println(entero_2)

	//* para cambiar el valor alojado en una direccion de memoria
	*puntero_1 = 7
	fmt.Println(entero_1)

}