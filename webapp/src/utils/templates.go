package utils

import (
	"fmt"
	"html/template"
	"net/http"
)

// Variavel padrao para se usar o Metodo Template do Go
var templates *template.Template

// Vai carregar todos os arquivos que estiverem dentro da pastas Views que tivem o .html
func CarregarTempletes() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// Essa funcao vai se auto execultar e vai renderizar uma pagina HTML
func ExecultarTemplate(w http.ResponseWriter, template string, dados interface{}) {
	templates.ExecuteTemplate(w, fmt.Sprintf("%s.html", template), dados)
}
