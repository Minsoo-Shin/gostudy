package jwt

import (
	"ggurugi/pkg/config"
	eu "ggurugi/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TokenValidationMinutes   = 60 * 5
	RefreshValidationMinutes = 60 * 60 * 24 * 365
	_id                      = "id"
	_role                    = "role"
	_username                = "username"
	_iat                     = "iat"
	_exp                     = "exp"
)

type CreateTokenRequest struct {
	ID       primitive.ObjectID
	UserName string
	UserRole string
	Duration time.Duration
}

func (ctr CreateTokenRequest) GetClaims() jwt.MapClaims {
	claims := jwt.MapClaims{
		_id:       ctr.ID,
		_role:     ctr.UserRole,
		_username: ctr.UserName,
		_iat:      time.Now().Unix(),
		_exp:      time.Now().Add(ctr.Duration).Unix(),
	}
	return claims
}

type Jwt interface {
	CreateToken(req CreateTokenRequest) (string, error)
}

type jwtToken struct {
	cfg config.Config
}

func New(cfg config.Config) Jwt {
	return &jwtToken{
		cfg: cfg,
	}
}

func (j *jwtToken) CreateToken(request CreateTokenRequest) (string, error) {
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, request.GetClaims()).SignedString([]byte(j.cfg.Jwt.Secret))
	if err != nil {
		return "", eu.InternalError(err)
	}
	return token, nil
}
