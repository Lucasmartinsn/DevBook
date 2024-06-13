package controllers

import "net/http"

func CriarUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuario"))
}
func BuscaUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuarios"))
}
func BuscaUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuarios"))
}
func AtualizarUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando o usuario"))
}
func DeletaUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando um usuario"))
}
