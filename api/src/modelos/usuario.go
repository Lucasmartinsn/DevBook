package modelos

import (
	"errors"
	"strings"
	"time"

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

func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validacao(etapa); err != nil {
		return err
	} else {
		usuario.formatar()
		return nil
	}
}

func (ususario *Usuario) validacao(etapa string) (erro error) {
	for _, key := range []string{ususario.Nome, ususario.Nick, ususario.Email} {
		if key == "" {
			erro = errors.New("Variavel requeridas estao vindo vazias, verifique o sua request JSON")
			return
		} else if ususario.Senha == "" && etapa == "cadastrar" {
			erro = errors.New("Senha nao deve ser vazia")
			return
		}
	}
	if erro = checkmail.ValidateFormat(ususario.Email); erro != nil {
		erro = errors.New("formato de email invalido")
		return
	}
	return nil
}

func (ususario *Usuario) formatar() {
	ususario.Nome = strings.TrimSpace(ususario.Nome)
	ususario.Nick = strings.TrimSpace(ususario.Nick)
	ususario.Email = strings.TrimSpace(ususario.Email)
}
