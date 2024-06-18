package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	Id       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEM time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {
	if err := usuario.validacao(); err != nil {
		return err
	} else {
		usuario.formatar()
		return nil
	}
}

func (ususario *Usuario) validacao() (erro error) {
	for _, key := range []string{ususario.Nome, ususario.Nick, ususario.Email, ususario.Senha} {
		if key == "" {
			erro = errors.New("Variavel requeridas estao vindo vazias, verifique o sua request JSON")
			return
		}
	}
	return nil
}

func (ususario *Usuario) formatar() {
	ususario.Nome = strings.TrimSpace(ususario.Nome)
	ususario.Nick = strings.TrimSpace(ususario.Nick)
	ususario.Email = strings.TrimSpace(ususario.Email)
}
