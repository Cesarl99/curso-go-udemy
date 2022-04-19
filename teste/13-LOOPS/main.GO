package main

import ("fmt"
)

func main(){
//	i := 0
/*
	for i < 10{
		i++
		fmt.Println("acrecentando no i ",i)
		//time.Sleep(time.Second)

	}
	fmt.Println(i)

	for j := 0 ;j < 10 ; j += 2{
		fmt.Println("acrecentando no j ",j)
		//time.Sleep(time.Second)
	}	
	*/
	//range
	nome:= [3]string{"joão","cesar","paulo"}
	for _ ,nome := range nome{
		fmt.Println(nome)
	}
	for indice,letra := range "PALAVRA"{
		fmt.Println(indice,string(letra))
	}
	usuario := map[string]string{
		"nome " : "cesar",
		"sobrenome" : "luiz",
	}
	for chave,valor := range usuario{
		fmt.Println(chave,valor)
	}
}