package test

import (
	"testing"

	"github.com/badrunp/media-sosial-web/be/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashAndComparePassword(t *testing.T) {

	var password string = "rahasia"
	hash := utils.HashAndSalt([]byte(password))
	match := utils.ComparePasswords(hash,[]byte("rahasia"))
	assert.Equal(t, true, match)

}

func TestHashAndComparePasswordFail(t *testing.T) {

	var password string = "rahasia"
	hash := utils.HashAndSalt([]byte(password))
	match := utils.ComparePasswords(hash,[]byte("tidakrahasia"))
	assert.NotEqual(t, true, match)

}