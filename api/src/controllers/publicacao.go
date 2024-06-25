package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/Lucasmartinsn/DevBook/api/src/autenticacao"
	"github.com/Lucasmartinsn/DevBook/api/src/banco"
	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
	"github.com/Lucasmartinsn/DevBook/api/src/repositorios"
	"github.com/Lucasmartinsn/DevBook/api/src/resposta"
	"github.com/gorilla/mux"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	bodyResquest, err := io.ReadAll(r.Body)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	var publicacao modelos.Publicacao
	if err = json.Unmarshal(bodyResquest, &publicacao); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	if err = publicacao.PrepararPublicacao("cadastrar"); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	reposositorio := repositorios.NewReporOfPublicacao(conn)
	publicacaoId, err := reposositorio.Criar(publicacao, userId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 201, fmt.Sprintf("Usuario cadastradro, ID: %d", publicacaoId))
}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	reposositorio := repositorios.NewReporOfPublicacao(conn)
	publicacao, err := reposositorio.BucarPublicacao(publicacaoId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, publicacao)
}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
