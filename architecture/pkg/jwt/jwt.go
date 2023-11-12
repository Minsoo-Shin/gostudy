package jwt

import (
	"github.com/Minsoo-Shin/go-boilerplate/pkg/config"
	eu "github.com/Minsoo-Shin/go-boilerplate/pkg/errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
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

type jwtToken struct {
	cfg config.Config
}

func New(cfg config.Config) Jwter {
	return &jwtToken{
		cfg: cfg,
	}
}

type TokenClaim struct {
	ID       string
	Duration time.Duration
}

func (j *jwtToken) NewToken(tokenClaim TokenClaim) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(tokenClaim.Duration)),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.cfg.Jwt.Secret))
	if err != nil {
		return "", eu.InternalError(err)
	}
	return token, nil
}

func (j *jwtToken) Verfiy(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

	switch {
	case token.Valid:
		return nil
	case errors.Is(err, jwt.ErrTokenMalformed):
		return ErrTokenMalformed
	case errors.Is(err, jwt.ErrTokenSignatureInvalid):
		return ErrTokenInvalidSignature
	case errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet):
		return ErrTokenExpired
	default:
		return ErrTokenInValid
	}
}
