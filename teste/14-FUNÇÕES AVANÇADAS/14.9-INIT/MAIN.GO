package main

import "fmt"

var n int


func main(){
	fmt.Println("escutando a função main")
	fmt.Println(n)
}

func init(){
	fmt.Println("executando a função init")
	n = 10
}