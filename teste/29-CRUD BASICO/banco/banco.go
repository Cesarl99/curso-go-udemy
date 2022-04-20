package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver de conexão
)

//Conectar abre a conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	StringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", StringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil
}
