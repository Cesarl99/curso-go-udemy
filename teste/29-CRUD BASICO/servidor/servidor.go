package servidor

import (
	"CRUD/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

//criarUsuarios e insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corporequisição, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("falha ao ler o corpo da requisição"))
		return
	}
	var usuario usuario

	if erro = json.Unmarshal(corporequisição, &usuario); erro != nil {
		w.Write([]byte("erro ao converter"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao converter conectar no banco de dados!! "))
		return
	}

	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome,email) values (?,?)")
	if erro != nil {
		w.Write([]byte("erro ao criar o statement!!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("erro ao executar o statement!!"))
		return
	}
	idinserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("erro ao executar o statement!"))
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuario inserido com sucesso ! Id : %d ", idinserido)))
}

//BuscaUsarios traz todos os usuarios salvos no bancos de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("ERRO AO CONECTAR COM O BANCO DE DADOS!"))
	}
	defer db.Close()
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("ERRO AO BUSCAR OS USUARIOS"))
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("erro ao escanerar o usuario"))
			return
		}
		usuarios = append(usuarios, usuario)
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Erro ao converter os usuarios para JSON"))
	}
}

//BuscaUsarios traz UM os usuarios salvos no bancos de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o paramentro para inteiro"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar ao banco de dados"))
		return
	}
	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("erro ao buscar o usuario"))
		return
	}
	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao buscar o usuario!!"))
			return
		}
	}
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("erro ao converter o usuario para JSON!"))
		return
	}
}

//AtualizaUsuario altera os dados de um usuários no banco
func AtualizarUsuarios(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("erro ao converter o parametro para inteiro"))
	}
	corporequisição, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("erro ao ler o corpo da requisição !"))
		return
	}
	var usuario usuario
	if erro := json.Unmarshal(corporequisição, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuario para struct"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar no banco de dados!!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?,email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("erro ao clicar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("erro ao atualizar o usuario!"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//DeleteUsuario remove um usuario do banco de dados
func DeletarUsuarios(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Erro ao converter o parametro para inteiro"))
		return
	}
	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("erro ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("erro ao criar o statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("erro ao deletar o usuario"))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
