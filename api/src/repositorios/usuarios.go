package repositorios

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statment, erro := repositorio.db.Prepare("insert into usuarios (nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}

	insercao, erro := statment.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := insercao.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Usuarios) BuscarUsuarios(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)
	linhas, erro := repositorio.db.Query("select id, nome, nick, email from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuariosBuscados []models.Usuario
	for linhas.Next() {
		var usuarioBuscado models.Usuario
		if erro = linhas.Scan(
			&usuarioBuscado.ID,
			&usuarioBuscado.Nome,
			&usuarioBuscado.Nick,
			&usuarioBuscado.Email,
		); erro != nil {
			return nil, erro
		}
		usuariosBuscados = append(usuariosBuscados, usuarioBuscado)
	}

	return usuariosBuscados, nil
}

func (usuario Usuarios) BusacarUsuarioPorID(usuarioID uint64) (models.Usuario, error) {
	var usuarioBuscado models.Usuario
	linha, erro := usuario.db.Query("select id, nome, nick, email from usuarios where id = ?", usuarioID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	if linha.Next() {
		erro = linha.Scan(
			&usuarioBuscado.ID,
			&usuarioBuscado.Nome,
			&usuarioBuscado.Nick,
			&usuarioBuscado.Email,
		)
		if erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuarioBuscado, nil
}

func (usuario Usuarios) BusacarUsuarioPorEmail(email string) (models.Usuario, error) {
	var usuarioBuscado models.Usuario
	linha, erro := usuario.db.Query("select id, senha from usuarios where email = ?", email)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linha.Close()

	if linha.Next() {
		if erro = linha.Scan(&usuarioBuscado.ID, &usuarioBuscado.Senha); erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuarioBuscado, nil
}

func (usuario Usuarios) AtualizarUsuario(usuarioID uint64, usuarioPassado models.Usuario) error {
	statment, erro := usuario.db.Prepare("update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}

	if _, erro = statment.Exec(usuarioPassado.Nome, usuarioPassado.Nick, usuarioPassado.Email, usuarioID); erro != nil {
		return erro
	}

	return nil
}

func (usuario Usuarios) DeletarUsuario(usuarioID uint64) error {
	statment, erro := usuario.db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(usuarioID); erro != nil {
		return erro
	}
	return nil
}
