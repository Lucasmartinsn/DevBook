package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/banco"
	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
	"github.com/Lucasmartinsn/DevBook/api/src/repositorios"
	"github.com/Lucasmartinsn/DevBook/api/src/resposta"
)

func CriarUser(w http.ResponseWriter, r *http.Request) {
	bodyResquest, err := io.ReadAll(r.Body)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	var usuario modelos.Usuario
	if err = json.Unmarshal(bodyResquest, &usuario); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	reposositorio := repositorios.NewReporOfUser(conn)
	userId, err := reposositorio.Criar(usuario)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 201, fmt.Sprintf("Usuario cadastradro, ID: %d", userId))
}

func BuscaUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}
func BuscaUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuarios"))
}
func AtualizarUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando o usuario"))
}
func DeletaUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuario"))
}
