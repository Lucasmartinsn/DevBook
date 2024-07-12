package controllers

import (
	"encoding/json"
	"errors"
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
	if err = publicacao.PrepararPublicacao(); err != nil {
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
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	reposositorio := repositorios.NewReporOfPublicacao(conn)
	publicacoes, err := reposositorio.BuscarPublicacoes(userId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, publicacoes)
}
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
	publicacao, err := reposositorio.BuscarPublicacao(publicacaoId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, publicacao)
}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfPublicacao(conn)
	publicacaoSalvaDB, err := repositorio.BuscarPublicacao(publicacaoId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	if publicacaoSalvaDB.AutorId != userId {
		resposta.Erro(w, 403, errors.New("nao é possivel atualizar um Post que nao seja seu"))
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
	if err = publicacao.PrepararPublicacao(); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	if err = repositorio.AtualizarPublicacoes(publicacaoId, publicacao); err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	resposta.Json(w, 200, nil)
}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfPublicacao(conn)
	publicacaoSalvaDB, err := repositorio.BuscarPublicacao(publicacaoId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	if publicacaoSalvaDB.AutorId != userId {
		resposta.Erro(w, 403, errors.New("nao é possivel remover um Post que nao seja seu"))
		return
	}
	if err = repositorio.Deletepublicacao(publicacaoId); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 204, nil)
}
func BuscarPublicacaoUser(w http.ResponseWriter, r *http.Request) {
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()
	repositorio := repositorios.NewReporOfPublicacao(conn)
	publicacoes, err := repositorio.BuscarPublicacoesUser(userId)
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, publicacoes)
}
func LikePublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	repositorio := repositorios.NewReporOfPublicacao(conn)
	if publicacao, err := repositorio.BuscarPublicacao(publicacaoId); err != nil {
		resposta.Erro(w, 500, err)
		return
	} else if publicacao.AutorId == userId {
		resposta.Erro(w, 403, errors.New("voce nao pode curtir suas proprias publicações"))
		return
	}

	if err = repositorio.Curtir(publicacaoId); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, nil)
}
func UnlikePublicacao(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	publicacaoId, err := strconv.ParseUint(param["id"], 10, 64)
	if err != nil {
		resposta.Erro(w, 400, err)
		return
	}
	userId, err := autenticacao.ExtrairID(r)
	if err != nil {
		resposta.Erro(w, 401, err)
		return
	}
	conn, err := banco.Connction()
	if err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	defer conn.Close()

	repositorio := repositorios.NewReporOfPublicacao(conn)
	if publicacao, err := repositorio.BuscarPublicacao(publicacaoId); err != nil {
		resposta.Erro(w, 500, err)
		return
	} else if publicacao.AutorId == userId {
		resposta.Erro(w, 403, errors.New("voce nao pode curtir suas proprias publicações"))
		return
	}

	if err = repositorio.DesCurtir(publicacaoId); err != nil {
		resposta.Erro(w, 500, err)
		return
	}
	resposta.Json(w, 200, nil)
}
