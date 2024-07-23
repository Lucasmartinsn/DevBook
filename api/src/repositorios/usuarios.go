package repositorios

import (
	"database/sql"
	"encoding/json"
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
		`select u.id, u.nome, u.nick, u.email, u.criacaoEm,
			COALESCE(
				(
					SELECT JSON_ARRAYAGG(JSON_OBJECT('id', u2.id, 'nome', u2.nome, 'nick', u2.nick))
					FROM usuario u2
					INNER JOIN seguidores s ON u2.id = s.usuarioId
					WHERE s.seguidoresId = u.id
				),
				JSON_ARRAY()
			) AS seguidores
		from usuario u
		where u.nome LIKE ? or nick LIKE ?;`, nomeOUnicks, nomeOUnicks,
	)
	if err != nil {
		return nil, err
	}
	defer linha.Close()

	var usuarios []modelos.Usuario
	for linha.Next() {
		var usuario modelos.Usuario
		var seguidoresJson string

		if err = linha.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriacaoEM , &seguidoresJson); err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(seguidoresJson), &usuario.Seguidores); err != nil {
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
func (repositorio usuario) BuscarPorEmail(email string) (modelos.Usuario, error) {
	linha, err := repositorio.db.Query("select id, senha from usuario where email = ?", email)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Id, &usuario.Senha); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}
func (repositorio usuario) Seguir(seguidorId, usuarioId uint64) error {
	statement, err := repositorio.db.Prepare("insert ignore into seguidores (usuarioId, seguidoresId) values (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioId, seguidorId); err != nil {
		return err
	}
	return nil
}
func (repositorio usuario) PararDeSeguir(seguidorId, usuarioId uint64) error {
	statement, err := repositorio.db.Prepare("delete from seguidores where usuarioId = ? and  seguidoresId = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuarioId, seguidorId); err != nil {
		return err
	}
	return nil
}
func (repositorio usuario) BuscarSeguidores(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(`select u.id, u.nome, u.nick, u.email, u.criacaoEm from usuario u inner join
		seguidores s on u.id = s.usuarioId where s.seguidoresId = ?`, usuarioId)
	if err != nil {
		return nil, err
	}
	var seguidores []modelos.Usuario
	for linhas.Next() {
		var seguidor modelos.Usuario
		if err = linhas.Scan(&seguidor.Id, &seguidor.Nome, &seguidor.Nick, &seguidor.Email, &seguidor.CriacaoEM); err != nil {
			return nil, err
		}
		seguidores = append(seguidores, seguidor)
	}
	return seguidores, nil
}
func (repositorio usuario) BuscarSeguindo(usuarioId uint64) ([]modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(`select u.id, u.nome, u.nick, u.email, u.criacaoEm from usuario u inner 
		join seguidores s on u.id = s.seguidoresId where s.usuarioId = ?`, usuarioId)
	if err != nil {
		return nil, err
	}
	var usuarios []modelos.Usuario
	for linhas.Next() {
		var usuario modelos.Usuario
		if err = linhas.Scan(&usuario.Id, &usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriacaoEM); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}
func (repositorio usuario) BuscarPorSenha(id uint64) (string, error) {
	linha, err := repositorio.db.Query("select senha from usuario where id = ?", id)
	if err != nil {
		return "", err
	}
	defer linha.Close()

	var usuario modelos.Usuario
	if linha.Next() {
		if err = linha.Scan(&usuario.Senha); err != nil {
			return "", err
		}
	}
	return usuario.Senha, nil
}
func (repositorio usuario) AtualizarSenha(id uint64, senha string) error {
	statement, err := repositorio.db.Prepare("update usuario set senha = ? where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(senha, id); err != nil {
		return err
	}
	return nil
}
