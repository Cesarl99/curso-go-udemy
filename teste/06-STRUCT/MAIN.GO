package main

import ("fmt")

type usuario struct{
	nome string
	idade int
	endereço endereço
}
type endereço struct{
	logadouro string
	numero int

}

func main(){
	fmt.Println("arquivo struct")

	var x usuario
	x.nome = "Cesar"
	x.idade = 22
	fmt.Println(x)

	EndereçoUsuario01 := endereço{"avenida São Paulo",252}
	usuario2 := usuario {"Lucas",20,EndereçoUsuario01}
	fmt.Println(usuario2)

	usuario3 := usuario {idade:30}
	fmt.Println(usuario3)
}