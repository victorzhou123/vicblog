package authimpl

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/victorzhou123/vicblog/common/domain/auth"
	cmdmerror "github.com/victorzhou123/vicblog/common/domain/error"
	cmprimitive "github.com/victorzhou123/vicblog/common/domain/primitive"
)

const (
	mapClaimsFieldNameUsername = "username"
	mapClaimsFieldNameExp      = "exp"
)

type timeAdder interface {
	AddUnix(time.Duration) int64
}

func NewSignJwt(timeAdder timeAdder, cfg *Config) *signJwt {
	return &signJwt{
		timeAdder:  timeAdder,
		secretKey:  cfg.SecretKey,
		expireTime: cfg.expireTime(),
	}
}

type signJwt struct {
	timeAdder
	secretKey  string
	expireTime time.Duration
}

func (s *signJwt) GenToken(pl *auth.Payload) (string, error) {
	return s.genSignBasedUsername(pl.UserName.Username())
}

func (s *signJwt) TokenValid(sign string) (cmprimitive.Username, error) {
	username, err := s.verifyToken(sign)
	if err != nil {
		return nil, err
	}

	return cmprimitive.NewUsername(username)
}

func (s *signJwt) genSignBasedUsername(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		mapClaimsFieldNameUsername: username,
		mapClaimsFieldNameExp:      s.AddUnix(s.expireTime), // set expire time at xxx
	})

	// generate sign jwt string by jwt.token
	return token.SignedString([]byte(s.secretKey))
}

func (s *signJwt) verifyToken(sign string) (string, error) {

	token, err := jwt.Parse(sign, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("parse sign error: %s", err.Error())
	}

	if !token.Valid {
		return "", fmt.Errorf("token invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("assert mapClaims failed")
	}

	if claims.VerifyNotBefore(int64(s.expireTime), true) {
		return "", cmdmerror.New(cmdmerror.ErrorCodeTokenInvalid, "")
	}

	username, ok := claims[mapClaimsFieldNameUsername].(string)
	if !ok {
		return "", cmdmerror.New(cmdmerror.ErrorCodeTokenInvalid, "")
	}

	return username, nil
}
