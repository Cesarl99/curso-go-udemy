package main

import ("fmt")

func função1(){
	fmt.Println("executando a função 1")
}

func funcao2(){
	fmt.Println("executando a função 2")
}
//DEFER = ADIAR
func AlunoEstaAprovado(n1,n2 float32)bool{
	defer fmt.Println("Media calculada.resultado será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1/n2)/2
	
	if media >= 6{
		return true
	}
	return false
}
func main(){
	fmt.Println(AlunoEstaAprovado(7,8))
}