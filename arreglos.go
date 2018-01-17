package main
 import "fmt"

func main() {
	//Al ser declarados de esta manera, los 10 lugares adquieren valor 0
	var arreglo_1 [10]int
	//Al ser declarados de esta manera, los 10 lugares adquieren valor ""
	var arreglo_2 [10]string
	//Al ser declarados de esta manera, los 10 lugares adquieren valor false
	var arreglo_3 [10]bool

	arreglo_4 :=[5]int{1,2,3,4,5}
	//Si no se entregan todos los valores, se completa con 0
	arreglo_5 :=[5]int{1,2,3}
	arreglo_5[3]=4

	var matriz_1 [2][2]int
	var matriz_2 [1][2]int
	var matriz_3 [2][3]int
	matriz_3[0][1] = 5
	
	fmt.Println(arreglo_1)
	fmt.Println(arreglo_2)
	fmt.Println(arreglo_3)
	fmt.Println(arreglo_4)
	fmt.Println(arreglo_5)

	//cantidad de elementos
	fmt.Println(len(arreglo_5))

	for i := 0; i < len(arreglo_5); i++ {
		fmt.Println(arreglo_5[i])
	}

	fmt.Println(matriz_1)
	fmt.Println(matriz_2)
	fmt.Println(matriz_3)
	fmt.Println(matriz_3[0][1])
}
