package utility

import (
	"errors"
	"healthcare-allopurinol/sentinel"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimsContent struct {
	Id         int64 `json:"id"`
	Role       int64 `json:"role"`
	IsVerified bool  `json:"is_verified"`
}

type customClaims struct {
	ClaimsContent
	jwt.RegisteredClaims
}

func GenerateJWToken(content ClaimsContent) (string, error) {
	now := time.Now()
	registeredClaims := customClaims{
		ClaimsContent: content,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: os.Getenv("APP_NAME"),
			IssuedAt: &jwt.NumericDate{
				Time: now,
			},
			ExpiresAt: &jwt.NumericDate{
				Time: now.Add(48 * time.Hour),
			},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)

	tokenString, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWToken(tokenString string) (*ClaimsContent, error) {
	var srcClaim customClaims

	tokenString = strings.Split(tokenString, "Bearer ")[1]
	token, err := jwt.ParseWithClaims(
		tokenString,
		&srcClaim,
		func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("TOKEN_SECRET")), nil
		},
		jwt.WithIssuer(os.Getenv("APP_NAME")),
		jwt.WithExpirationRequired(),
		jwt.WithValidMethods([]string{"HS256", "HS384", "HS512"}),
	)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, sentinel.ErrTokenExpired
		}
		return nil, err
	}

	if !token.Valid {
		return nil, sentinel.ErrTokenInvalid
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return nil, errors.New("unexpected claim type")
	}

	return &claims.ClaimsContent, nil
}
