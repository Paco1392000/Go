//estructura construidas
//su base es un arraylog
//es dinamico
//los arreglos ya tienen un tamaño definido no camviable
//no define su tammaño en el tiempo de ejecución

package main

import (
"fmt"

)

func main(){
	//ya que no tiene el valor es un slice
	var matriz[]int{1,2,3,4}
	fmt.Println(matriz)
	//al inicializar lso arreglso, todos los elementos del arreglo es 0
	//en el caso de los slices su valor inicial es nulo
	var matrizvacia[]int
	if matrizvacia == nil{
		fmt.Println("esta vacio")
	}else{
		fmt.Println(len(matriz))//len contiene longitud
	}
//puntero al arreglo
//longitud del arreglo al que apunta
//capacidad
	
	//otra forma de obtener un slice es atravez de un entero
	arreglo := [3]int{1,2,3}
	//convertirlo a sclice
	slice := arreglo[:2]//parte un arreglo :n que del inicio a esa posición se obtenen los datos
	// n:  de n posición al valor final se obtienen los valors

}