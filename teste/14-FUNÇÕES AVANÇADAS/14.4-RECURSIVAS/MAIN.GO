package main

import ("fmt")

func fibonacci(posição uint) uint{
	if posição <=1{
		return posição
	}
	return fibonacci(posição - 2)+fibonacci(posição -1)
}
func main(){
	fmt.Println("funções recursivas")

	posição := uint(100)
	for i:=uint(1);i<= posição;i++{
		fmt.Println(fibonacci(i))
	}

}	