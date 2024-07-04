package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/service"
)

// Vai ser usada para acresentar o token na resquet
func FazerRequestWithAuth(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(metodo, url, dados)
	if err != nil {
		return nil, err
	}
	cookie, _ := service.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"])

	cliet := &http.Client{}
	response, err := cliet.Do(request)
	if err != nil {
		return nil, err
	}
	return response, err
}
