package main

import( 
"fmt"
"strconv"
)

func main(){
	///---------------------------------
	//--- 1 --- Declaracion de variables
	//----------------------------------
	w:=1// declaracion denominada tipado dinamico
	// el compilador en tiempo de ejecucion decide que tipo es
	var x,y,z int = 2,3,4// valor inicial 0
	var palabra string = "Hola"// valor inicial cadena vacia
	var bandera bool = true// valor inicial false
	cadenas :=[]string {"palabra 1","palabra 2"}
	//Go no compilara hasta que las variables sean utilizadas

	//--- 1 --- salida
	fmt.Println("-- Salida de la unidad 1 --")
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(palabra)
	fmt.Println(bandera)
	fmt.Println(cadenas)

	///---------------------------------
	//--- 2 --- Conversion de variables
	//----------------------------------
	var variable_1 int = 17
	//Convertir un Int -> string
	var variable_2 string = strconv.Itoa(variable_1)
	//Convertir un string -> Int 
	//El operador _ nos ayuda a recibir variables que no requerimos usar
	variable_3, _ := strconv.Atoi(variable_2)

	fmt.Println("-- Salida de la unidad 2 --")
	fmt.Println(variable_1)
	fmt.Println("El valor de la variable_1 es de: "+variable_2)
	fmt.Println(variable_3+1)

	///---------------------------------
	//--- 5 --- Ciclo
	//----------------------------------
	fmt.Println("-- Salida de la unidad 5 --")
	fmt.Println("-- For tipo 1:")
	for variable_4:=0 ; variable_4<10 ; variable_4++ {
		fmt.Printf("valor de variable_4 :%d\n",variable_4)
	}
	fmt.Println("-- For tipo 2:")
	variable_5 := 0
	for variable_5<10{
		fmt.Printf("valor de variable_5 :%d\n",variable_5)
		variable_5++
	}
	fmt.Println("-- For tipo 3:")
	variable_6 := 0
	for {
		fmt.Printf("valor de variable_6 :%d\n",variable_6)
		variable_6++
		if variable_6 > 10{
			break
		}
	}
	fmt.Println("-- For tipo 4:")
	variable_7 := 0
	for {
		if variable_7 == 3{
			variable_7++
			continue //corta el for hasta aca, y lo reinicia conservando el valor de variable_7
		}
		fmt.Printf("valor de variable_7 :%d\n",variable_7)
		variable_7++
		if variable_7 > 10{
			break
		}
	}
}