package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyProvider(t *testing.T) {
	testService := NewOpenIdAuthService()
	config := ""

	_, err := testService.Authenticate(config)

	assert.NotNil(t, err)
}

func TestProviderWithValue(t *testing.T) {
	testService := NewOpenIdAuthService()
	config := "https://maxgit.us.auth0.com/"
	authenticator, err := testService.Authenticate(config)

	assert.NotNil(t, authenticator)
	assert.Nil(t, err)
}

func TestInitSession(t *testing.T) {
	testService := NewOpenIdAuthService()
	err := testService.InitSession()
	assert.Nil(t, err)
}

func TestGetStore(t *testing.T) {
	testService := NewOpenIdAuthService()
	_ = testService.InitSession()
	store := testService.NewStore()
	assert.NotNil(t, store)
}
