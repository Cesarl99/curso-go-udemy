package main

import ("fmt")

func main(){
	fmt.Println("estrutura de controle")

	numero:=10

	if numero > 15 {
		fmt.Println("maior que 15")
	}else{
		fmt.Println("menor ou igual a 15")
	}
	if outroNumero := numero; outroNumero > 0{
		fmt.Println("numero é maior que zero")
	}else{
		fmt.Println("numero é menor que zero")
	}
}