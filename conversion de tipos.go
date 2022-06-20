package main
//se importa un paquete para luego ejecutar el programa
//se pueden usar ; pero no es necesario
//en la ventana para ejecutar, se usa la palabra go build nombreaplicacion.go
import "fmt"
func main(){
	edad := 22
	//go no conversiona los enteros a strings
	edad_str := strconv.Itoa(edad)
	//en el caso se al convertir edad_int,_(o err) := strconv.Atoi(variable)
	fmt.Println("edad de: "+edad_str)
	//Convertir string a entero strconv.Atoi(variable)

}