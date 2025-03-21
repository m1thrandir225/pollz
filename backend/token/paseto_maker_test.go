package token

import (
	"m1thrandir225/cicd2025/util"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomPasetoMaker(t *testing.T) TokenMaker {
	testKey := util.RandomString(32)

	tokenMaker, err := NewPasetoMaker(testKey)

	require.NoError(t, err)
	require.NotEmpty(t, tokenMaker)
	return tokenMaker
}

func createToken(t *testing.T) (TokenMaker, string) {
	userId := uuid.New()
	maker := createRandomPasetoMaker(t)

	tokenExpirationTime := time.Minute * 15

	token, payload, err := maker.CreateToken(userId, tokenExpirationTime)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotEmpty(t, token)

	require.Equal(t, payload.UserId, userId)
	require.WithinDuration(t, payload.ExpiredAt, time.Now().Add(tokenExpirationTime), time.Second)

	return maker, token
}

func TestNewPasetoMaker(t *testing.T) {
	createRandomPasetoMaker(t)
}

func TestCreateToken(t *testing.T) {
	createToken(t)
}

func TestVerifyToken(t *testing.T) {
	maker, token := createToken(t)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
}
