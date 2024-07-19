package token

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/saidamir98/udevs_pkg/logger"
)

type JWTHandler struct {
	UserId    string
	Exp       string
	Iot       string
	Aud       []string
	Role      string
	SignedKey string
	Log       logger.LoggerI
	Token     string
	TimeOut   int
}

type CostumClaims struct {
	*jwt.Token
	UserId string   `json:"user_id"`
	Exp    int      `json:"exp"`
	Iot    string   `json:"iot"`
	Role   string   `json:"user_role"`
	Aud    []string `json:"aud"`
}

func (j *JWTHandler) GenerateToken() (accesToken string, err error) {

	var (
		token *jwt.Token
		claim jwt.MapClaims
	)

	token = jwt.New(jwt.SigningMethodHS256)

	claim = token.Claims.(jwt.MapClaims)
	claim["user_id"] = j.UserId
	claim["exp"] = time.Now().Add(time.Minute * time.Duration(j.TimeOut)).Unix()
	claim["user_role"] = j.Role

	accesToken, err = token.SignedString([]byte(j.SignedKey))
	if err != nil {
		j.Log.Error("error on Generating token", logger.Error(err))
		return
	}
	return
}
