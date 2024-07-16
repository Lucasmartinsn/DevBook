package service

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
)

var S *securecookie.SecureCookie

func ConfigCookie() {
	S = securecookie.New([]byte(os.Getenv("HASHKEY")), []byte(os.Getenv("BLOCKKEY")))
}

func Salvar(w http.ResponseWriter, id, token string) error {
	dados := map[string]string{
		"id":    id,
		"token": token,
	}
	dadosCodificados, err := S.Encode("dados", dados)
	if err != nil {
		return err
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",
		HttpOnly: true,
	})
	return nil
}

// Var verificar de o usuario tem o Cookie
func Ler(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("dados")
	if err != nil {
		return nil, err
	}
	valores := make(map[string]string)
	if err = S.Decode("dados", cookie.Value, &valores); err != nil {
		return nil, err
	}

	return valores, nil
}

// Essa função vai apagar o Cookie
func Delete(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})
}
