package modelos

import (
	"errors"
	"strings"
	"time"

	"github.com/Lucasmartinsn/DevBook/api/src/seguranca"
	"github.com/badoux/checkmail"
)

type Usuario struct {
	Id        uint64    `json:"id,omitempty"`
	Nome      string    `json:"nome,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Senha     string    `json:"senha,omitempty"`
	CriacaoEM time.Time `json:"criacaoEm,omitempty"`
}

type Senha struct {
	Atual string `json:"atual"`
	Nova  string `json:"nova"`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validacao(etapa); err != nil {
		return err
	} else {
		if err := usuario.formatar(etapa); err != nil {
			return err
		}
		return nil
	}
}

func resposta(keys []string) error {
	for _, key := range keys {
		if key == "" {
			return errors.New("Variavel requeridas estao vindo vazias, verifique o sua request JSON")
		}
	}
	return nil
}

func (ususario *Usuario) validacao(etapa string) (erro error) {
	switch etapa {
	case "cadastrar":
		if erro = resposta([]string{ususario.Nome, ususario.Nick, ususario.Email, ususario.Senha}); erro != nil {
			return erro
		}

	case "login":
		if erro = resposta([]string{ususario.Email, ususario.Senha}); erro != nil {
			return erro
		}

	case "atualizar":
		if erro = resposta([]string{ususario.Nome, ususario.Nick, ususario.Email}); erro != nil {
			return erro
		}
	}
	if erro = checkmail.ValidateFormat(ususario.Email); erro != nil {
		erro = errors.New("formato de email invalido")
		return
	}
	return nil
}

func (ususario *Usuario) formatar(etapa string) error {
	if etapa == "cadastrar" {
		senha, err := seguranca.Hash(ususario.Senha)
		if err != nil {
			return err
		}
		ususario.Senha = string(senha)
	}
	ususario.Nome = strings.TrimSpace(ususario.Nome)
	ususario.Nick = strings.TrimSpace(ususario.Nick)
	ususario.Email = strings.TrimSpace(ususario.Email)
	return nil
}
