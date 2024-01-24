package token

import (
	"testing"
	"time"

	"github.com/Lisciowsky/my_simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	symetricKey := []byte(util.RandomString(32))
	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	maker, err := NewPasetoMaker(symetricKey)
	require.NoError(t, err)

	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)

	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)

	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)

	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

func TestExpiredPasetoTokenTest(t *testing.T) {
	symetricKey := []byte(util.RandomString(32))
	username := util.RandomOwner()

	maker, err := NewPasetoMaker(symetricKey)
	require.NoError(t, err)

	token, err := maker.CreateToken(username, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
