package goauth

import (
	"encoding/json"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

type GoAuth interface {
	CreateToken(data map[string]interface{}, expiration time.Duration) (string, error)
	DecryptToken(tokenString string) (map[string]interface{}, error)
}

type goAuthImpl struct {
	SecretKey string
}

func NewGoAuth(secretKey string) GoAuth {
	return goAuthImpl{
		SecretKey: secretKey,
	}
}

func (auth goAuthImpl) CreateToken(data map[string]interface{}, expiration time.Duration) (string, error) {
	data["exp"] = time.Now().Add(expiration).Unix()

	claimMap := make(jwt.MapClaims)
	for k, v := range data {
		claimMap[k] = v
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claimMap)
	return at.SignedString([]byte(auth.SecretKey))
}

func (auth goAuthImpl) DecryptToken(tokenString string) (map[string]interface{}, error) {
	result, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(auth.SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return auth.ConvertToMap(result.Claims), nil
}

func (auth goAuthImpl) ConvertToMap(claim jwt.Claims) map[string]interface{} {
	claimMap := make(map[string]interface{})
	jsonString, _ := json.Marshal(claim)
	json.Unmarshal(jsonString, &claimMap)
	return claimMap
}
