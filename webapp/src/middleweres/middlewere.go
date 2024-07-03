package middleweres

import (
	"fmt"
	"net/http"
	"webapp/src/service"
)

// Vai criar os logs das requisições
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// vai verificar se o usuario esta logado
func Autenticacao(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := service.Ler(r); err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next(w, r)
	}
}
