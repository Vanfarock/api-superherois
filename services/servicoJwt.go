package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type IServicoJWT interface {
	GerarToken(usuario string) string
	ValidarToken(token string) (*jwt.Token, error)
}

type claimsUsuario struct {
	Nome string `json:"nome"`
	jwt.StandardClaims
}

type ServicoJWT struct {
	chaveSecreta string
	emissor      string
}

func ObterServicoAutenticacaoJWT() IServicoJWT {
	return ServicoJWT{
		chaveSecreta: getChaveSecreta(),
	}
}

func getChaveSecreta() string {
	segredo := os.Getenv("ISSO_NAO_EH_IMPORTANTE")
	if segredo == "" {
		return "VN+x+rw5nhB5OMh0EV/t4g=="
	}
	return segredo
}

func (servico ServicoJWT) GerarToken(usuario string) string {
	doisDias := time.Hour * 48
	claims := claimsUsuario{
		usuario,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(doisDias).Unix(),
			Issuer:    servico.emissor,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenAssinado, err := token.SignedString([]byte(servico.chaveSecreta))
	if err != nil {
		panic(err)
	}
	return tokenAssinado
}

func (service ServicoJWT) ValidarToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, service.parseToken)
}

func (service ServicoJWT) parseToken(token *jwt.Token) (interface{}, error) {
	if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
		return nil, fmt.Errorf("token inválido %v", token.Header["alg"])
	}

	return []byte(service.chaveSecreta), nil
}
