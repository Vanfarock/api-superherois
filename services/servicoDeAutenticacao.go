package services

import (
	"prog-web/dao"
	"prog-web/database"
	"prog-web/models"
)

type ServicoDeAutenticacao struct{}

func (ServicoDeAutenticacao) Login(usuario, senha string) (bool, error) {
	db, err := database.Connect()
	if err != nil {
		return false, err
	}

	credenciais := models.Credenciais{
		Usuario: usuario,
		Senha:   senha,
	}

	autenticacaoDAO := new(dao.AutenticacaoDAO)
	return autenticacaoDAO.Login(db, credenciais), nil
}

func (ServicoDeAutenticacao) Cadastrar(usuario, senha string) error {
	db, err := database.Connect()
	if err != nil {
		return err
	}

	credenciais := models.Credenciais{
		Usuario: usuario,
		Senha:   senha,
	}

	autenticacaoDAO := new(dao.AutenticacaoDAO)
	return autenticacaoDAO.Cadastrar(db, &credenciais)
}
