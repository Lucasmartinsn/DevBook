package models

type DatosAuth struct {
	Token  string `json:"token"`
	IdUser uint64 `json:"idUser"`
}
