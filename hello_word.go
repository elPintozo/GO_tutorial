package main

import "fmt"
func main(){
	//--- 1 --- Declaracion de variables

	w:=1// declaracion denominada tipado dinamico
	// el compilador en tiempo de ejecucion decide que tipo es
	var x,y,z int = 2,3,4// valor inicial 0
	var palabra string = "Hola"// valor inicial cadena vacia
	var bandera bool = true// valor inicial false
	cadenas :=[]string {"palabra 1","palabra 2"}
	//Go no compilara hasta que las variables sean utilizadas

	//--- 1 --- salida
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(palabra)
	fmt.Println(bandera)
	fmt.Println(cadenas)
}