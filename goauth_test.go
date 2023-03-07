package goauth

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTokenWithOutExpiration(t *testing.T) {
	// A way to create a new instance of the assert library.
	assert := assert.New(t)
	goAuthInstance := NewGoAuth("my-secret-key")

	data := make(map[string]interface{})
	data["value1"] = "Mama soy paquito"
	data["value2"] = 4587

	const TimeOutToken = 3 * time.Second
	token, err := goAuthInstance.CreateToken(data, TimeOutToken)

	assert.NotEqual("", token)
	assert.Equal(true, strings.Contains(token, "ey"))
	assert.Nil(err)

	claim, errDecrypt := goAuthInstance.DecryptToken(token)
	assert.Nil(errDecrypt)
	assert.Equal("Mama soy paquito", claim["value1"])
	assert.Equal(float64(4587), claim["value2"])

}

func TestCreateTokenWithExpiration(t *testing.T) {
	assert := assert.New(t)
	goAuthInstance := NewGoAuth("my-secret-key")

	data := make(map[string]interface{})
	data["value1"] = "Mama soy paquito"
	data["value2"] = 4587

	const TimeOutToken = 1 * time.Second
	token, err := goAuthInstance.CreateToken(data, TimeOutToken)

	assert.NotEqual("", token)
	assert.Equal(true, strings.Contains(token, "ey"))
	assert.Nil(err)

	time.Sleep(2 * time.Second)
	claim, errDecrypt := goAuthInstance.DecryptToken(token)
	assert.Equal(0, len(claim))
	assert.NotNil(errDecrypt)
	assert.Error(errors.New("Token is expired"), errDecrypt)
}

func TestDecryptTokenFail(t *testing.T) {
	assert := assert.New(t)
	goAuthInstance := NewGoAuth("my-secret-key")
	claim, errDecrypt := goAuthInstance.DecryptToken("this_is_a_not_valid_token")
	assert.Equal(0, len(claim))
	assert.NotNil(errDecrypt)
	assert.Equal("token contains an invalid number of segments", errDecrypt.Error())
}
