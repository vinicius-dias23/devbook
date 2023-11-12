package controllers

import (
	"api/src/autenticacao"
	"api/src/banco"
	"api/src/models"
	"api/src/repositorios"
	"api/src/respostas"
	"api/src/seguranca"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Logar(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuarioBuscadoRequisicao models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuarioBuscadoRequisicao); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioUsuarios(db)
	var usuarioBuscadoBanco models.Usuario
	usuarioBuscadoBanco, erro = repositorio.BusacarUsuarioPorEmail(usuarioBuscadoRequisicao.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioBuscadoBanco.Senha, usuarioBuscadoRequisicao.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioBuscadoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
	}

	w.Write([]byte(token))
}
