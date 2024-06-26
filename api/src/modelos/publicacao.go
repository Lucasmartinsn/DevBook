package modelos

import (
	"errors"
	"strings"
	"time"
)

type Publicacao struct {
	Id        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorid,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

func (publicacao *Publicacao) PrepararPublicacao() error {
	if err := publicacao.validacaoPublicacao(); err != nil {
		return err
	} else {
		if err := publicacao.formatarPublicacao(); err != nil {
			return err
		}
		return nil
	}
}

func (publicacao *Publicacao) validacaoPublicacao() (erro error) {
	keys := []string{publicacao.Titulo, publicacao.Conteudo}
	for _, key := range keys {
		if key == "" {
			return errors.New("Variavel requeridas estao vindo vazias, verifique o sua request JSON")
		}
	}
	return nil
}

func (publicacao *Publicacao) formatarPublicacao() error {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
	return nil
}
