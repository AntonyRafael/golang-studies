package server

import (
	"db/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct"))
		return
	}

	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuarios (nome, email) values (?, ?)")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement!"))
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		w.Write([]byte("Falha ao executar a insercao!"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		w.Write([]byte("Falha ao obter o id inserido!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

// BuscarUsuarios lista todos os usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()

	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		w.Write([]byte("Falha ao buscar os usuários!"))
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Falha ao escanear o usuário!"))
			return
		}

		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		w.Write([]byte("Falha ao converter os usuários para JSON!"))
		return
	}
}

// BuscarUsuario lista um usuário específico
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco de dados!"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("Falha ao buscar o usuário!"))
		return
	}

	var usuario usuario
	if linha.Next() {
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Falha ao escanear o usuário!"))
			return
		}
	}

	if usuario.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado!"))
		return
	}

	if erro := json.NewEncoder(w).Encode(usuario); erro != nil {
		w.Write([]byte("Falha ao converter usuário para JSON!"))
		return
	}
}

// AtualizarUsuario atualiza um usuário específico
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter o parâmetro para inteiro!"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro := json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Falha converter usuário!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement!"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil {
		w.Write([]byte("Falha ao executar a atualização!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Usuário atualizado com sucesso!"))
}

// DeletarUsuario deleta um usuário específico
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("Falha ao converter o parâmetro para inteiro!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Falha ao conectar no banco de dados!"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("Falha ao criar statement!"))
		return
	}

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("Falha ao executar a deleção!"))
		return
	}
	defer statement.Close()

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("Usuário deletado com sucesso!"))
}
