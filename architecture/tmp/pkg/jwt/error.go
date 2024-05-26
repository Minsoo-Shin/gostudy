package jwt

import "github.com/pkg/errors"

var (
	ErrTokenInValid          = errors.New("ERR_TOKEN_INVALID")
	ErrTokenMalformed        = errors.New("ERR_TOKEN_MALFORMED")
	ErrTokenInvalidSignature = errors.New("ERR_INVALID_SIGNATURE")
	ErrTokenExpired          = errors.New("ERR_TOKEN_EXPIRED")
)
