package repositorios

import (
	"database/sql"
	"fmt"

	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
)

type usuario struct {
	db *sql.DB
}

func NewReporOfUser(db *sql.DB) *usuario {
	return &usuario{db}
}
func (repositorio usuario) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare("insert into usuario (nome, nick, email, senha) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}
	lastInsert, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsert), nil
}
func (repositorio usuario) Buscar(nomeOUnick string) ([]modelos.Usuario, error) {
	nomeOUnicks := fmt.Sprintf("%%%s%%", nomeOUnick)
	linha, err := repositorio.db.Query(
		"select id, nome,nick, email, criacaoEm from usuario where nome LIKE ? or nick LIKE ?", nomeOUnicks, nomeOUnicks,
	)
	if err != nil {
		return nil, err
	}
	defer linha.Close()

	var usuarios []modelos.Usuario
	for linha.Next() {
		var usuario modelos.Usuario
		if err = linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriacaoEM); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
func (repositorio usuario) BuscaUser(id uint64) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("select id, nome,nick, email, criacaoEm from usuario where id = ?", id)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriacaoEM); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}
func (repositorio usuario) Atualizar(id uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare("update usuario set nome = ?, nick = ?, email = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); err != nil {
		return err
	}
	return nil
}

func (repositorio usuario) Deletar(id uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuario where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(id); err != nil {
		return err
	}
	return nil
}