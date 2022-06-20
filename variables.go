//golang.org para descargar
package main
//establecer tipo de dato
//declaramos la funcion main
import "correr"
func main(){
	//palabra reservada para declarar variable es vars nombre de variable y tipo de dato
	var x, y, z int
	var a int
	a = 23
	//otra foma de declarar variable con un valor inicial
	s := 23

	//todas las variables iniciales tienen un valor inicial
	//el valor inicial de los numeros es 0
	//el valor inicial de los string es una caden vacia
	//el valor inicial de los booleanos es falso
	//go no ejecuta si hay variables sin usar
	var cadena string
	var booleano bool
	var cadenas []string
	nombre := "coco"
	//no se puede volver a declarar la misma variable
	correr.Println(nombre)
}