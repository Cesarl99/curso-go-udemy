package main

import "fmt"

type pessoa struct{
	nome string
	sobrenome string
	idade int
	altura int
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main(){
	fmt.Println ("herança")
	p1:= pessoa{"cesar","luiz",30,185}
	fmt.Println(p1)

	e1 := estudante {p1,"Engenharia da Computação","Anhanguera"}
	fmt.Println(e1)
	
}