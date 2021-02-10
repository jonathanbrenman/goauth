package goauth

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestCreateTokenWithOutExpiration(t *testing.T) {
	assert := assert.New(t)
	goAuth := NewGoAuth("my-secret-key")

	data := make(map[string]interface{})
	data["hola"] = "Mundo"
	data["pepe"] = 4587

	token, err := goAuth.CreateToken(data, 0)
	assert.NotEqual("", token)
	assert.Equal(true, strings.Contains(token, "ey"))
	assert.Nil(err)

	claim, err := goAuth.DecryptToken(token)
	assert.Nil(err)

	assert.Equal("Mundo", claim["hola"])
	assert.Equal(float64(4587), claim["pepe"])
}

func TestCreateTokenWithExpiration(t *testing.T) {
	assert := assert.New(t)
	goAuth := NewGoAuth("my-secret-key")

	data := make(map[string]interface{})
	data["myProperty1"] = "MyValue1"
	data["myProperty2"] = 10.5

	token, err := goAuth.CreateToken(data, 10 * time.Minute)
	assert.NotEqual("", token)
	assert.Equal(true, strings.Contains(token, "ey"))
	assert.Nil(err)

	claim, err := goAuth.DecryptToken(token)
	assert.Nil(err)

	assert.Equal("MyValue1", claim["myProperty1"])
	assert.Equal(10.5, claim["myProperty2"])
}

func TestDecryptTokenFail(t *testing.T) {
	assert := assert.New(t)
	goAuth := NewGoAuth("my-secret-key")
	_, err := goAuth.DecryptToken("this_is_a_not_valid_token")
	assert.Equal("token contains an invalid number of segments", err.Error())
}