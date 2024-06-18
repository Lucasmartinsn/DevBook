package repositorios

import (
	"database/sql"

	"github.com/Lucasmartinsn/DevBook/api/src/modelos"
)

type usuario struct {
	db *sql.DB
}

func NewReporOfUser(db *sql.DB) *usuario {
	return &usuario{db}
}

func (reposositorio usuario) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := reposositorio.db.Prepare("insert into usuario (nome, nick, email, senha) values (?,?,?,?)")
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
