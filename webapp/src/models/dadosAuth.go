package models

type DadosAuth struct {
	Token  string `json:"token"`
	IdUser uint64 `json:"idUser"`
}
