package autenticacao

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(id uint64) (string, error) {
	permissao := jwt.MapClaims{}
	permissao["authorized"] = true
	permissao["exp"] = time.Now().Add(time.Hour * 3).Unix()
	permissao["usuarioId"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissao)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

