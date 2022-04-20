package main

import (
	"database/sql"
	"fmt"
	"log"

	_"github.com/go-sql-driver/mysql"
)

func main() {
	StringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", StringConexao)
	if erro != nil {
		fmt.Println("dentro do SQL.Open")
		log.Fatal(erro)
	}
	defer db.Close()

	if erro = db.Ping();erro!=nil{
		fmt.Println("dentro do ping")
		log.Fatal(erro)
	}
	fmt.Println("conexão está aberta!!")

	linhas,erro := db.Query("select * from usuarios")
		if erro != nil{
			log.Fatal(erro)
		}
	defer linhas.Close()

	fmt.Println(linhas)
}
