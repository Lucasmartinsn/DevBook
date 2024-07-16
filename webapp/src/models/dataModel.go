package models

import "time"

type DadosAuth struct {
	Token  string `json:"token"`
	IdUser uint64 `json:"idUser"`
}

type Publicacao struct {
	Id        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorId   uint64    `json:"autorid,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

type Usuario struct {
	Id          uint64       `json:"id,omitempty"`
	Nome        string       `json:"nome,omitempty"`
	Nick        string       `json:"nick,omitempty"`
	Email       string       `json:"email,omitempty"`
	CriacaoEM   time.Time    `json:"criacaoEm,omitempty"`
	Seguidores  []Usuario    `json:"seguidores,omitempty"`
	Seguindo    []Usuario    `json:"seguindo,omitempty"`
	Publicacoes []Publicacao `json:"publicacoes,omitempty"`
}

type Senha struct {
	Atual string `json:"atual"`
	Nova  string `json:"nova"`
}
