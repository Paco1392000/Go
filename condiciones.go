package main

import(){
	"fmt"
}
/*
	== igual a :evalua que ambas expresiones sean iguales
	!= diferente
	<
	>
	<=
	>=
	&& operador AND cuando todas las expresiones sean verdaderas
	|| operador OR que una expresion sea verdadera
*/
func main(){
	x := 10
	y := 5
	//condiciones a ver
	if false{
		fmt.Println("holamundo")
	}
	if x > 5 {//los parentesis son opcionales
	//las llaves para la condicional no van en la linea siguiente del condicional
		fmt.Prinf("%d es mayor que %d \n",x,y)

	}else if x < y{//si no se ejecuta la anterior ejecutar esta
		fmt.Println("entre el else y el if")
	}else{
		tmt.Println("Entra al else")
	}



}