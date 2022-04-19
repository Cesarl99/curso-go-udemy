package main

import ("fmt"
		"mod-teste/enderecos")

func main(){
	TipoEndereços := endereços.TipodeEndereços("Avenida Paulista")
	fmt.Println(TipoEndereços)
}