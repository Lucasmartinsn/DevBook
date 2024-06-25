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

func (publicacao *Publicacao) PrepararPublicacao(etapa string) error {
	if err := publicacao.validacaoPublicacao(etapa); err != nil {
		return err
	} else {
		if err := publicacao.formatarPublicacao(); err != nil {
			return err
		}
		return nil
	}
}
func respostaPublicacao(keys []string) error {
	for _, key := range keys {
		if key == "" {
			return errors.New("Variavel requeridas estao vindo vazias, verifique o sua request JSON")
		}
	}
	return nil
}

func (publicacao *Publicacao) validacaoPublicacao(etapa string) (erro error) {
	switch etapa {
	case "cadastrar":
		if erro = resposta([]string{publicacao.Titulo, publicacao.Conteudo}); erro != nil {
			return erro
		}
	case "atualizar":
		if erro = resposta([]string{publicacao.Titulo, publicacao.Conteudo}); erro != nil {
			return erro
		}
	}
	return nil
}
func (publicacao *Publicacao) formatarPublicacao() error {
	publicacao.Titulo = strings.TrimSpace(publicacao.Titulo)
	publicacao.Conteudo = strings.TrimSpace(publicacao.Conteudo)
	return nil
}
