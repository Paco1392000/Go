package main

import(
	"fmt"
	"bufio"//utiliza reader
	"os"
)

func main(){

	edad := 2

	//leer datos de la termina
	var ingreso int
	fmt.Println("colocar valor: ")
	fmt.Scanf("%d\n",&ingreso)
	//se puede poner variagles antes de imprimir %d, %v, %f, %b, %s
	fmt.Println("mi edad es; %d\n", edad)
	fmt.Printf("valor ingresado %d\n", ingreso)
	//para string
	bandera := true
	fmt.Printf("FSD %t", bandera)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ingresar nombre: ")
	text,err := reader.ReadString("\n")
	reader.ReadString("\n")
	if (err != nil){//nulo en go es nil
		fmt.Println(err)
	}else{
		fmt.Println("Hola "+text)
	}



}