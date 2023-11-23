package repositorios

import (
	"api/src/models"
	"database/sql"
)

type Publicacoes struct {
	db *sql.DB
}

func NovoRepositorioPublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

func (repositorio Publicacoes) Criar(publicacao models.Publicacao) (uint64, error) {
	statment, erro := repositorio.db.Prepare("insert into publicacoes (titulo, conteudo, autor_id) values (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statment.Close()

	resultado, erro := statment.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Publicacoes) BuscarPublicacao(publicacaoID uint64) (models.Publicacao, error) {
	linha, erro := repositorio.db.Query("select p.*, u.nick from publicacoes p inner join usuarios u on p.autor_id = u.id where p.id = ?", publicacaoID)
	if erro != nil {
		return models.Publicacao{}, erro
	}
	defer linha.Close()

	var publicacaoBuscada models.Publicacao
	if linha.Next() {
		if erro = linha.Scan(
			&publicacaoBuscada.ID,
			&publicacaoBuscada.Titulo,
			&publicacaoBuscada.Conteudo,
			&publicacaoBuscada.AutorID,
			&publicacaoBuscada.Curtidas,
			&publicacaoBuscada.CriadaEm,
			&publicacaoBuscada.AutorNick,
		); erro != nil {
			return models.Publicacao{}, erro
		}
	}

	return publicacaoBuscada, nil
}

func (repositorio Publicacoes) BuscarPublicacoes(usuarioID uint64) ([]models.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`select distinct p.*, u.nick from publicacoes p
		inner join usuarios u on u.id = p.autor_id
		inner join seguidores s on p.autor_id = s.usuario_id
		where u.id = ? or s.seguidor_id = ?
		order by 1 desc`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoesBuscadas []models.Publicacao
	for linhas.Next() {
		var publicacaoBuscada models.Publicacao
		if erro = linhas.Scan(
			&publicacaoBuscada.ID,
			&publicacaoBuscada.Titulo,
			&publicacaoBuscada.Conteudo,
			&publicacaoBuscada.AutorID,
			&publicacaoBuscada.Curtidas,
			&publicacaoBuscada.CriadaEm,
			&publicacaoBuscada.AutorNick,
		); erro != nil {
			return nil, erro
		}

		publicacoesBuscadas = append(publicacoesBuscadas, publicacaoBuscada)
	}

	return publicacoesBuscadas, nil
}
