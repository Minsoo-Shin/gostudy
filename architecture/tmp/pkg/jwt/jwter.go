package jwt

type Jwter interface {
	NewToken(req TokenClaim) (string, error)
}
