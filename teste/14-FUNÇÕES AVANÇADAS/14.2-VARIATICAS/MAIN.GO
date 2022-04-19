package main

import ("fmt")

func soma(numeros ...int)int{
	total:=0
	for _,numero := range numeros{
		total += numero
	}
	return total
}
func escrever (texto string,numeros ... int){
	for _,numero:=range numeros{
		fmt.Println(texto,numero)
	}
}
func main(){
	totalSoma:=soma(2,3,32,5,5,4,7,8,5,24,14)
	fmt.Println(totalSoma)

	escrever("ola,mundo",20,36,74,100,121)
	
}