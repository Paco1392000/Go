package main

import(
	"fmt"

)

func main(){
	//go tiene for
	//for iniciador; booleana; i++{} 
	for i:= 0; i<10; i++{
		fmt.Println("hello")
		fmt.Println(i)
	}
	for i:= 0; i<10; i=i+2{
		fmt.Println(i)
	}
	//para simular ciclo while
	for i<10{
		fmt.Println(i)
		i++
	}
	//ciclo infinito
	i := 0
	for{
		if i==2{
			i++
			continue//continue vuelve al inicio del ciclo
		}

		fmt.Println(i)
		i++
		if i>10{
			break
		}
	}


}