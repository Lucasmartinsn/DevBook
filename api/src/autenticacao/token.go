package autenticacao

import (
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	// jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(id uint64) (string, error) {
	permissao := jwt.MapClaims{}
	permissao["authorized"] = true
	permissao["exp"] = time.Now().Add(time.Hour * 3).Unix()
	permissao["usuarioId"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissao)
	// return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, returnKeyToken)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token invalido")
}

func returnKeyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, errors.New("metodo de assinatura inesperado")
	}
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}
