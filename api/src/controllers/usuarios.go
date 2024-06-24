package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/Lucasmartinsn/DevBook/api/src/autenticacao"
	"github.com/Lucasmartinsn/DevBook/api/src/banco"
	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
	"github.com/Lucasmartinsn/DevBook/api/src/repositorios"
	"github.com/Lucasmartinsn/DevBook/api/src/resposta"
	"github.com/Lucasmartinsn/DevBook/api/src/seguranca"
	"github.com/gorilla/mux"
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
	if err := usuario.Preparar("cadastrar"); err != nil {
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
	userId, err := reposositorio.Criar(usuario)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 201, fmt.Sprintf("Usuario cadastradro, ID: %d", userId))
}
func BuscaUsers(w http.ResponseWriter, r *http.Request) {
	nameORnick := strings.ToLower(r.URL.Query().Get("usuario"))
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	repositorio := repositorios.NewReporOfUser(conn)
	usuarios, err := repositorio.Buscar(nameORnick)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, usuarios)
}
func BuscaUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(param["usuario"], 10, 64)
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

	repositorio := repositorios.NewReporOfUser(conn)
	user, err := repositorio.BuscaUser(usuarioId)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	resposta.Json(w, 200, user)
}
func AtualizarUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	if userId, erro := autenticacao.ExtrairID(r); erro != nil {
		resposta.Erro(w, 401, err)
		return
	} else if userId != usuarioId {
		resposta.Erro(w, 403, errors.New("sem permição para modificar usuarios que nao sejam o seu"))
		return
	}

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
	if err := usuario.Preparar("atualizar"); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfUser(conn)
	if err = repositorio.Atualizar(usuarioId, usuario); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, nil)
}
func DeletaUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	if userId, erro := autenticacao.ExtrairID(r); erro != nil {
		resposta.Erro(w, 401, err)
		return
	} else if userId != usuarioId {
		resposta.Erro(w, 403, errors.New("sem permição para remover usuarios que nao seja o seu"))
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfUser(conn)
	if err = repositorio.Deletar(usuarioId); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, nil)
}
func SeguirUser(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	seguidorId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	userId, erro := autenticacao.ExtrairID(r)
	if erro != nil {
		resposta.Erro(w, 400, err)
		return
	} else if userId == seguidorId {
		resposta.Erro(w, 403, errors.New("voce nao pode serguir voce mesmo"))
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfUser(conn)
	if err = repositorio.Seguir(seguidorId, userId); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, nil)
}
func PararDeSeguir(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	seguidorId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	userId, erro := autenticacao.ExtrairID(r)
	if erro != nil {
		resposta.Erro(w, 400, err)
		return
	} else if userId == seguidorId {
		resposta.Erro(w, 403, errors.New("nao é possivel parar de serguir voce mesmo"))
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	repositorio := repositorios.NewReporOfUser(conn)
	if err = repositorio.PararDeSeguir(seguidorId, userId); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, nil)
}
func BuscarSeguidores(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(param["id"], 10, 64)
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

	repositorio := repositorios.NewReporOfUser(conn)
	seguidores, err := repositorio.BuscarSeguidores(usuarioId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, seguidores)
}
func BuscarSeguindo(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(param["id"], 10, 64)
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

	repositorio := repositorios.NewReporOfUser(conn)
	seguidores, err := repositorio.BuscarSeguindo(usuarioId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, seguidores)
}
func AtualizarUserPass(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	usuarioId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	if userId, erro := autenticacao.ExtrairID(r); erro != nil {
		resposta.Erro(w, 400, err)
		return
	} else if userId != usuarioId {
		resposta.Erro(w, 403, errors.New("voce nao pode modificar um usuario que nao seja o seu"))
		return
	}
	bodyResquest, err := io.ReadAll(r.Body)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	var senha modelos.Senha
	if err = json.Unmarshal(bodyResquest, &senha); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfUser(conn)
	senhaAtual, err := repositorio.BuscarPorSenha(usuarioId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	if err := seguranca.VerificaSenha(senhaAtual, senha.Atual); err != nil {
		resposta.Erro(w, 401, errors.New("senha atual nao coresponde a senha salva no banco"))
		return
	}
	newSenha, err := seguranca.Hash(senha.Nova)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	if err := repositorio.AtualizarSenha(usuarioId, string(newSenha)); err != nil {
		resposta.Erro(w, 500, err)
		return
	}

	resposta.Json(w, 200, nil)
}
