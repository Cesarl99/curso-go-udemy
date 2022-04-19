package main

import ("fmt")

func main(){

	retorno := func (texto string)string{
		return fmt.Sprintf("recebido -> %s",texto)
	}("mais um teste")
	fmt.Println(retorno)
}