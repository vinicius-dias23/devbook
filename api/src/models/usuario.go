package models

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) PrepararUsuario(etapa string) error {
	if erro := usuario.validarCampos(etapa); erro != nil {
		return erro
	}

	usuario.formatarCampos()
	return nil
}

func (usuario *Usuario) validarCampos(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo nome não pode estar vazio!")
	}
	if usuario.Nick == "" {
		return errors.New("O nick nome não pode estar vazio!")
	}
	if usuario.Email == "" {
		return errors.New("O campo e-mail não pode estar vazio!")
	}
	if etapa == "create" && usuario.Senha == "" {
		return errors.New("O campo Senha não pode estar vazio!")
	}

	return nil
}

func (usuario *Usuario) formatarCampos() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
