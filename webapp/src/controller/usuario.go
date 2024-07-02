package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Vai chamar a API para a criaçao de um novo Cadastro de usuario
func CadastroOfUser(w http.ResponseWriter, r *http.Request) {
	// Ele vai pegar o Bode da request e vai deixar os dados acessiveis para manipulação
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(usuario)
}
