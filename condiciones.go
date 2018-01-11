package main

import "fmt"

func main() {

	variable_1 := 5
	variable_2 := 10
	variable_3 := 2
	///---------------------------------
	//--- 4 --- Condiciones
	//----------------------------------
	/*
		== 
		!=
		<
		>
		>=
		<=
		&& ->  and
		|| -> or
	*/
	fmt.Println("-- Salida de la unidad 4 --")
	if true{
		fmt.Println("Verdad")
	}
	if variable_1 > variable_2 {
		fmt.Printf("%d es mayor que %d \n",variable_1,variable_2)
	}else{
		fmt.Printf("%d es es menor que %d \n",variable_1,variable_2)
	}

	if variable_3 == 1{
		fmt.Printf("Seleccionaste la opcion %d \n",variable_3)
	}else if variable_3 == 2 {
		fmt.Printf("Seleccionaste la opcion %d y no la 1\n",variable_3)	
	}else{
		fmt.Printf("Seleccionaste la opcion %d, y las opciones eran 1 o 2 \n",variable_3)
	}
}