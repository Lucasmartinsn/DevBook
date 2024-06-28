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
func (repositorio publicacao) BuscarPublicacao(id uint64) (modelos.Publicacao, error) {
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
func (repositorio publicacao) BuscarPublicacoes(id uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`select distinct p.*, u.nick from publicacao p 
		inner join usuario u on u.id = p.autorId 
		inner join seguidores s on p.autorId = s.usuarioId where p.id = ? or s.seguidoresId = ? order by 1 desc`, id, id)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err = linhas.Scan(&publicacao.Id, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorId,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}
func (repositorio publicacao) AtualizarPublicacoes(id uint64, publicacao modelos.Publicacao) error {
	statement, err := repositorio.db.Prepare("update publicacao set titulo = ?, conteudo = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(publicacao.Titulo, publicacao.Conteudo, id); err != nil {
		return err
	}
	return nil
}
func (repositorio publicacao) Deletepublicacao(id uint64) error {
	statement, err := repositorio.db.Prepare("delete from publicacao where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}
func (repositorio publicacao) BuscarPublicacoesUser(id uint64) ([]modelos.Publicacao, error) {
	linhas, err := repositorio.db.Query(`
		select p.*, u.nick from publicacao p 
		join usuario u on u.id = p.autorId where p.autorId = ? order by 1 desc`, id)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if err = linhas.Scan(&publicacao.Id, &publicacao.Titulo, &publicacao.Conteudo, &publicacao.AutorId,
			&publicacao.Curtidas, &publicacao.CriadaEm, &publicacao.AutorNick); err != nil {
			return nil, err
		}
		publicacoes = append(publicacoes, publicacao)
	}
	return publicacoes, nil
}
func (repositorio publicacao) Curtir(id uint64) error {
	statement, err := repositorio.db.Prepare("update publicacao set curtidas = curtidas + 1 where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}
func (repositorio publicacao) DesCurtir(id uint64) error {
	statement, err := repositorio.db.Prepare(`
		update publicacao set curtidas = CASE 
		WHEN curtidas > 0 THEN curtidas - 1 
		ELSE curtidas END where id = ?`)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}
