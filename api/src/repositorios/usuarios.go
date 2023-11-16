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

func (usuario Usuarios) Seguir(usuarioId, seguidorId uint64) error {
	statment, erro := usuario.db.Prepare("insert ignore into seguidores (usuario_id, seguidor_id) values (?, ?)")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (usuario Usuarios) PararDeSeguir(usuarioId, seguidorId uint64) error {
	statment, erro := usuario.db.Prepare("delete from seguidores where usuario_id = ? and seguidor_id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(usuarioId, seguidorId); erro != nil {
		return erro
	}

	return nil
}

func (usuario Usuarios) BuscarSeguidores(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := usuario.db.Query("select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.seguidor_id where s.usuario_id = ?", usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var seguidoresBuscados []models.Usuario
	for linhas.Next() {
		var seguidorBuscado models.Usuario
		if erro = linhas.Scan(
			&seguidorBuscado.ID,
			&seguidorBuscado.Nome,
			&seguidorBuscado.Nick,
			&seguidorBuscado.Email,
			&seguidorBuscado.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		seguidoresBuscados = append(seguidoresBuscados, seguidorBuscado)
	}

	return seguidoresBuscados, nil
}

func (usuario Usuarios) BuscarSeguindo(usuarioId uint64) ([]models.Usuario, error) {
	linhas, erro := usuario.db.Query("select u.id, u.nome, u.nick, u.email, u.criadoEm from usuarios u inner join seguidores s on u.id = s.usuario_id where s.seguidor_id = ?", usuarioId)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var listaPessoasSeguindo []models.Usuario
	for linhas.Next() {
		var pessoaSeguindo models.Usuario
		if erro = linhas.Scan(
			&pessoaSeguindo.ID,
			&pessoaSeguindo.Nome,
			&pessoaSeguindo.Nick,
			&pessoaSeguindo.Email,
			&pessoaSeguindo.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		listaPessoasSeguindo = append(listaPessoasSeguindo, pessoaSeguindo)
	}

	return listaPessoasSeguindo, nil
}

func (usuario Usuarios) BuscarSenha(usuarioId uint64) (string, error) {
	linhas, erro := usuario.db.Query("select senha from usuarios where id = ?", usuarioId)
	if erro != nil {
		return "", nil
	}
	defer linhas.Close()

	var usuarioBuscado models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuarioBuscado.Senha); erro != nil {
			return "", erro
		}
	}

	return usuarioBuscado.Senha, nil
}

func (usuario Usuarios) AtualizarSenha(senhaNovaComHash string, usuarioId uint64) error {
	statment, erro := usuario.db.Prepare("update usuarios set senha = ? where id = ?")
	if erro != nil {
		return erro
	}
	defer statment.Close()

	if _, erro = statment.Exec(senhaNovaComHash, usuarioId); erro != nil {
		return erro
	}

	return nil
}
