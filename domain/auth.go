package domain

import (
	"database/sql"
	"mini-project/config"
	"mini-project/errs"
	"mini-project/logger"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// 24 hour
const ACCESS_TOKEN_DURATION = 24 * time.Hour

type Login struct {
	Username   string        `db:"username"`
	CustomerId sql.NullInt32 `db:"customer_id"`
	Role       string        `db:"role"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AccessTokenClaims struct {
	CustomerId sql.NullInt32 `json:"customer_id"`
	Username   string        `json:"username"`
	Role       string        `json:"role"`
	jwt.RegisteredClaims
}

type AuthToken struct {
	token *jwt.Token
}

func (l Login) ClaimsAccessToken() AccessTokenClaims {
	if l.CustomerId.Valid {
		return l.claimsForUser()
	} else {
		return l.claimsForAdmin()
	}
}

func (l Login) claimsForUser() AccessTokenClaims {
	return AccessTokenClaims{
		Username:   l.Username,
		Role:       l.Role,
		CustomerId: sql.NullInt32{Int32: l.CustomerId.Int32, Valid: true},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}
}

func (l Login) claimsForAdmin() AccessTokenClaims {
	return AccessTokenClaims{
		Username: l.Username,
		Role:     l.Role,
	}
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func (t AuthToken) NewAccessToken() (string, *errs.AppErr) {
	signedString, err := t.token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		logger.Error("Failed while signing access token: " + err.Error())
		return "", errs.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}
