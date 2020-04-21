package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()

	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, at.AccessToken, "", "new access token should not have a defined access token id")
	assert.True(t, at.UserId == 0, "new access token should not have a defined user id")
}

func TestAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token created three hours from now should not be expired")
}

func TestAccessTokenDuration(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}
