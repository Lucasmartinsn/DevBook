package repositorios

import (
	"database/sql"

	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
)

type publicacao struct {
	db *sql.DB
}

func NewReporOfPublicacao(db *sql.DB) *publicacao {
	return &publicacao{db}
}
func (repositorio publicacao) Criar(publicacao modelos.Publicacao, idUser uint64) (uint64, error) {
	statement, err := repositorio.db.Prepare("insert into publicacao (titulo, conteudo, autorId) values (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, idUser)
	if err != nil {
		return 0, err
	}
	lastInsert, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsert), nil
}
func (repositorio publicacao) BucarPublicacao(id uint64) (modelos.Publicacao, error) {
	linha, err := repositorio.db.Query(`select p.*, u.nick from publicacao p inner join usuario u on u.id = p.autorId where p.id = ?`, id)
	if err != nil {
		return modelos.Publicacao{}, err
	}
	defer linha.Close()

	var publicacao modelos.Publicacao
	if linha.Next() {
		if err = linha.Scan(&publicacao.Id, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorId,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return modelos.Publicacao{}, err
		}
	}
	return publicacao, nil
}
