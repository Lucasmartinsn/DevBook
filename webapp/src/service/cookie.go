package service

import (
	"net/http"
	"os"

	"github.com/gorilla/securecookie"
)

var S *securecookie.SecureCookie

func ConfigCookie() {
	S = securecookie.New([]byte(os.Getenv("HASHKEY")), []byte(os.Getenv("BLOCKKEY")))
}

func Salvar(w http.ResponseWriter, id uint64, token string) error {
	dados := map[string]any{
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
func Ler(r *http.Request) (map[string]any, error) {
	cookie, err := r.Cookie("dados")
	if err != nil {
		return nil, err
	}
	valores := make(map[string]any)
	if err = S.Decode("dados", cookie.Value, &valores); err != nil {
		return nil, err
	}

	return valores, nil
}
