package jwt

import (
	"net/http"
	"os"
	"strings"
	"time"

	"golang_example/model"

	"github.com/golang-jwt/jwt/v5"
)

var (
	signingKey = []byte(os.Getenv("SECRET_KEY"))
	issuer     = "golang_example"
	duration   = time.Duration(24 * time.Hour)
)

func GenerateAuthorizationToken(request model.AuthorizationRequest) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,
		"exp": jwt.NewNumericDate(time.Now().Add(duration)),
		"iat": jwt.NewNumericDate(time.Now()),
		"subject": map[string]any{
			"deivce":   request.Device,
			"platform": request.Platform,
			"language": request.Language,
		},
	})

	token, e := t.SignedString(signingKey)

	if e != nil {
		return "", e
	}

	return token, nil
}

func ValidateRequest(r *http.Request) error {
	token := r.Header.Get("Authorization")

	claims, e := ParseToken(token)

	if e != nil {
		return e
	}

	e = VerifyInfoToken(claims)

	if e != nil {
		return e
	}

	return nil
}

func ParseToken(token string) (jwt.Claims, error) {
	if strings.HasPrefix(token, "Bear") {
		token = strings.Replace(token, "Bear ", "", 1)
	}

	t, e := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method == jwt.SigningMethodHS256 {
			return signingKey, nil
		}

		return nil, jwt.ErrInvalidKey
	})

	if e != nil {
		return nil, e
	}

	return t.Claims, nil
}

func VerifyInfoToken(claims jwt.Claims) error {
	iss, e := claims.GetIssuer()

	if e != nil {
		return e
	}

	if strings.Compare(iss, issuer) != 0 {
		return jwt.ErrTokenInvalidIssuer
	}

	exp, e := claims.GetExpirationTime()

	if e != nil {
		return e
	}

	if exp.Before(time.Now()) {
		return jwt.ErrTokenExpired
	}

	iat, e := claims.GetIssuedAt()

	if e != nil {
		return e
	}

	if exp.Sub(iat.Time) != duration {
		return jwt.ErrTokenInvalidClaims
	}

	return nil
}
