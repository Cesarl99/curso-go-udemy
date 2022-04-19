package main

import "fmt"

func recuperarExecucao(){
	if r := recover();r != nil{
		fmt.Println("execução recuperada com  sucesso")
	}
}
func AlunoEstaAprovado1(n1,n2 float64) bool{
	defer recuperarExecucao()
	media := (n1+n2)/2

	if media > 6{
		return true
	}else if media < 6{
		return false
	}

	panic("A MEDIA É EXTAMENTE 6!!")
	
}

func main(){
	fmt.Println(AlunoEstaAprovado1(8,4))
	fmt.Println("PÓS EXECUÇÃO!!!")
}