package middlewares

import (
	"fmt"
	"net/http"

	"github.com/Lucasmartinsn/DevBook/api/src/autenticacao"
	"github.com/Lucasmartinsn/DevBook/api/src/resposta"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Sprintf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Sprintf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

func Autenticar(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := autenticacao.ValidarToken(r); err != nil {
			resposta.Erro(w, 401, err)
			return
		}
		next(w, r)
	}
}
