package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/banco"
	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
	"github.com/Lucasmartinsn/DevBook/api/src/repositorios"
)

func CriarUser(w http.ResponseWriter, r *http.Request) {
	bodyResquest, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var usuario modelos.Usuario
	if err = json.Unmarshal(bodyResquest, &usuario); err != nil {
		log.Fatal(err)
	}
	conn, err := banco.Connction()
	if err != nil {
		log.Fatal(err)
	}
	reposositorio := repositorios.NewReporOfUser(conn)
	if err != nil {
		log.Fatal(err)
	}
	userId, err := reposositorio.Criar(usuario)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Usuario cadastradro, ID: %d", userId)))
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
