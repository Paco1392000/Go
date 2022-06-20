package main

import(
	"fmt"
)

func main(){
	//tipo de arreglos
	//arreglo estatico se establece al compilar el archivo
		//var nombre del arreglo [cantidad] tipo
	var arreglo [10] int

	arreglo2 := [3]int {1,2,3}
	arreglo3 := [3]int{1,2}
	arreglo3[2]=20//en la tercera posición del arreglo se agrega el valor
	fmt.prinln(arreglo)
	//posición de los elementos en un arreglo es n-1
	//ejemplo de ello es el inicial =1 en un arreglo es 0

	for i:=0; i < len(arreglo); i++{
	
		fmt.Println(arreglo[i])

	}

	var matriz [3][2]int
	matriz[0][1] := 5

}