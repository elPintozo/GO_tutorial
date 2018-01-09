package main

import( 
"fmt"
"bufio"
"os"
)

func main(){
	///---------------------------------
	//--- 3 --- Imprimir y leer desde la terminal
	//----------------------------------
	var variable_1 int = 10
	var variable_2 bool = true
	var variable_3 string = "Ricardo"
	var variable_4 float32 = 1.55
	var variable_5 string = "Rancagua"

	fmt.Println("-- Salida de la unidad 3 --")
	//Imprimir variables en terminal   
	fmt.Printf("El int es: %d\n", variable_1)
	fmt.Printf("El booleano es: %t\n", variable_2)
	fmt.Printf("El string es: %v\n", variable_3)
	fmt.Printf("El float es: %f\n", variable_4)
	fmt.Printf("El string es: %s\n", variable_5)

	//Leer desde terminal
	var variable_6 int
	var variable_7 string

	fmt.Print("Tu edad es: ")
	fmt.Scanf("%d\n", &variable_6)
	fmt.Printf("Tienes: %d\n",variable_6)

	fmt.Print("Tu nombre es: ")
	fmt.Scanf("%s\n", &variable_7)
	fmt.Printf("Hola, %s\n",variable_7)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa tu nombre: ")
	text,err := reader.ReadString('\n')

	if err != nil{
		fmt.Print(err)
	}else{
		fmt.Printf("Bienvenido %s", text)
	}

}