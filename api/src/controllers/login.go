package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/banco"
	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
	"github.com/Lucasmartinsn/DevBook/api/src/repositorios"
	"github.com/Lucasmartinsn/DevBook/api/src/resposta"
	"github.com/Lucasmartinsn/DevBook/api/src/seguranca"
)

func Login(w http.ResponseWriter, r *http.Request) {
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
	if err := usuario.Preparar("login"); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	reposositorio := repositorios.NewReporOfUser(conn)
	data, err := reposositorio.BuscarPorEmail(usuario.Email)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	if err = seguranca.VerificaSenha(data.Senha, usuario.Senha); err != nil {
		resposta.Erro(w, 403, err)
		return
	}

	resposta.Json(w, 201, "token")
}
