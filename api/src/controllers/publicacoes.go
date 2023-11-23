package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarPublicacao(w http.ResponseWriter, r *http.Request) {
	usuarioID, erro := autenticacao.ExtrairUsuarioId(r)
	if erro != nil {
		respostas.JSON(w, http.StatusUnauthorized, erro)
		return
	}

	corpoDaRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publicacaoInserida models.Publicacao
	publicacaoInserida.AutorID = usuarioID

	if erro = json.Unmarshal(corpoDaRequisicao, &publicacaoInserida); erro != nil {
		respostas.JSON(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publicacaoInserida.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	publicacao := repositorios.NovoRepositorioPublicacoes(db)
	publicacaoInserida.ID, erro = publicacao.Criar(publicacaoInserida)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	respostas.JSON(w, http.StatusCreated, publicacaoInserida)
}
func BuscarPublicacoes(w http.ResponseWriter, r *http.Request) {

}
func BuscarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func AtualizarPublicacao(w http.ResponseWriter, r *http.Request) {

}
func DeletarPublicacao(w http.ResponseWriter, r *http.Request) {

}
