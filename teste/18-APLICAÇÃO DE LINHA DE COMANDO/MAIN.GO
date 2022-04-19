package main

import (
		"os"
		"log"
		"linha-de-comando/app"
	)

func main(){
	aplicacao := app.Gerar()
	if erro:=aplicacao.Run(os.Args);erro != nil{
		log.Fatal(erro)
	}

}
