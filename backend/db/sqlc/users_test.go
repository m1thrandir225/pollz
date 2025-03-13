package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"m1thrandir225/cicd2025/util"
	"testing"
)

func createRandomUser(t *testing.T) CreateUserRow {
	password := util.RandomString(16)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	firstName := util.RandomString(10)
	lastName := util.RandomString(10)
	email := util.RandomEmail()

	args := CreateUserParams{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hashedPassword,
	}

	user, err := testStore.CreateUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.NotZero(t, user.ID)
	require.NotEmpty(t, user.CreatedAt)
	require.Equal(t, user.FirstName, firstName)
	require.Equal(t, user.LastName, lastName)
	require.Equal(t, user.Email, email)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	result, err := testStore.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)

	require.Equal(t, user.CreatedAt, result.CreatedAt)
	require.Equal(t, user.Email, result.Email)
	require.Equal(t, user.ID, result.ID)
	require.Equal(t, user.FirstName, result.FirstName)
	require.Equal(t, user.LastName, result.LastName)
}

func TestGetUserDetails(t *testing.T) {
	user := createRandomUser(t)

	result, err := testStore.GetUserDetails(context.Background(), user.ID)

	require.NoError(t, err)

	require.Equal(t, user.CreatedAt, result.CreatedAt)
	require.Equal(t, user.Email, result.Email)
	require.Equal(t, user.ID, result.ID)
	require.Equal(t, user.FirstName, result.FirstName)
	require.Equal(t, user.LastName, result.LastName)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	result, err := testStore.DeleteUser(context.Background(), user.ID)

	require.NoError(t, err)

	require.Equal(t, user.CreatedAt, result.CreatedAt)
	require.Equal(t, user.Email, result.Email)
	require.Equal(t, user.ID, result.ID)
	require.Equal(t, user.FirstName, result.FirstName)
	require.Equal(t, user.LastName, result.LastName)

	resultDeleted, err := testStore.GetUserDetails(context.Background(), user.ID)
	require.Error(t, err)
	require.Empty(t, resultDeleted)
}
