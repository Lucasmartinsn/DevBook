package requisicoes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"webapp/src/models"
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

func buscarDadosOfUser(canal chan<- models.Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d", os.Getenv("BASE_URL"), usuarioId)
	response, err := FazerRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- models.Usuario{}
		return
	}
	defer response.Body.Close()

	var usuario models.Usuario
	if err = json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		canal <- models.Usuario{}
		return
	}
	canal <- usuario

}
func buscarDadosOfUserSeguidores(canal chan<- []models.Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d/seguidores", os.Getenv("BASE_URL"), usuarioId)
	response, err := FazerRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []models.Usuario{}
		return
	}
	defer response.Body.Close()

	var seguidores []models.Usuario
	if err = json.NewDecoder(response.Body).Decode(&seguidores); err != nil {
		canal <- []models.Usuario{}
		return
	}
	canal <- seguidores

}
func buscarDadosOfUserSeguidor(canal chan<- []models.Usuario, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/usuario/%d/seguindo", os.Getenv("BASE_URL"), usuarioId)
	response, err := FazerRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []models.Usuario{}
		return
	}
	defer response.Body.Close()

	var seguindo []models.Usuario
	if err = json.NewDecoder(response.Body).Decode(&seguindo); err != nil {
		canal <- []models.Usuario{}
		return
	}
	canal <- seguindo

}
func buscarDadosOfUserPublicacao(canal chan<- []models.Publicacao, usuarioId uint64, r *http.Request) {
	url := fmt.Sprintf("%s/publicacao/%d/usuario", os.Getenv("BASE_URL"), usuarioId)
	response, err := FazerRequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		canal <- []models.Publicacao{}
		return
	}
	defer response.Body.Close()

	var usuario []models.Publicacao
	if err = json.NewDecoder(response.Body).Decode(&usuario); err != nil {
		canal <- []models.Publicacao{}
		return
	}
	canal <- usuario

}

// Vai fazer 4 request na API para montar uma Struc de usuario completo
func BuscarUserFullWithAuth(id uint64, r *http.Request) (models.Usuario, error) {
	// Criando os canais para manipular as Go routines
	canalUsuario := make(chan models.Usuario)
	canalSeguidores := make(chan []models.Usuario)
	canalSeguidor := make(chan []models.Usuario)
	canalPublicacao := make(chan []models.Publicacao)

	// Chamando as funções encadeadas como Go routines
	go buscarDadosOfUser(canalUsuario, id, r)
	go buscarDadosOfUserSeguidores(canalSeguidores, id, r)
	go buscarDadosOfUserSeguidor(canalSeguidor, id, r)
	go buscarDadosOfUserPublicacao(canalPublicacao, id, r)

	// Variavei para serem usuadas No Select
	var (
		usuario     models.Usuario
		seguidores  []models.Usuario
		seguindo    []models.Usuario
		publicacoes []models.Publicacao
	)

	// Esse for vai monitorar os canas das Go routines e para retornar o struct Usuario completo
	for i := 0; i < 4; i++ {
		select {
		case usuarioReload := <-canalUsuario:
			if usuarioReload.Id == 0 {
				// return models.Usuario{}, errors.New("erro ao buscar o Usuario")
				usuario = models.Usuario{}
			}
			usuario = usuarioReload

		case seguidoresReload := <-canalSeguidores:
			if seguidoresReload == nil {
				// return models.Usuario{}, errors.New("erro ao buscar os Seguidores")
				seguidores = []models.Usuario{}
			}
			seguidores = seguidoresReload

		case seguindoReload := <-canalSeguidor:
			if seguindoReload == nil {
				// return models.Usuario{}, errors.New("erro ao buscar os usuarios Seguindo")
				seguindo = []models.Usuario{}
			}
			seguindo = seguindoReload

		case publicacaoReload := <-canalPublicacao:
			if publicacaoReload == nil {
				// return models.Usuario{}, errors.New("erro ao buscar as publicações")
				publicacoes = []models.Publicacao{}
			}
			publicacoes = publicacaoReload

		}

	}
	usuario.Seguindo = seguindo
	usuario.Seguidores = seguidores
	usuario.Publicacoes = publicacoes

	return usuario, nil
}
