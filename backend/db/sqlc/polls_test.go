package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"m1thrandir225/cicd2025/util"
	"testing"
)

func createRandomPoll(t *testing.T, userId uuid.UUID) CreatePollRow {
	pollDescription := util.RandomString(120)

	args := CreatePollParams{
		Description: pollDescription,
		CreatedBy:   userId,
	}
	poll, err := testStore.CreatePoll(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, poll)
	require.Equal(t, pollDescription, poll.Description)
	require.Equal(t, userId, poll.CreatedBy)
	require.NotZero(t, poll.CreatedAt)

	return poll
}

func TestCreatePoll(t *testing.T) {
	user := createRandomUser(t)
	createRandomPoll(t, user.ID)
}

func TestDeletePoll(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	result, err := testStore.DeletePoll(context.Background(), poll.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, poll.ID, result.ID)
	require.Equal(t, user.ID, poll.CreatedBy)
	require.NotZero(t, poll.CreatedAt)

	deleted, err := testStore.GetPoll(context.Background(), poll.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}

func TestGetPoll(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	result, err := testStore.GetPoll(context.Background(), poll.ID)

	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, poll.ID, result.ID)
	require.Equal(t, poll.CreatedBy, result.CreatedBy)
	require.Equal(t, poll.CreatedAt, result.CreatedAt)
	require.Equal(t, poll.Description, result.Description)
}

func TestUpdatePoll(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	newDescription := util.RandomString(120)

	updateArgs := UpdatePollParams{
		ID:          poll.ID,
		Description: newDescription,
		IsActive:    true,
	}
	result, err := testStore.UpdatePoll(context.Background(), updateArgs)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, poll.ID, result.ID)
	require.Equal(t, newDescription, result.Description)
	require.NotEqual(t, poll.Description, result.Description)
	require.Equal(t, poll.IsActive, result.IsActive)
	require.Equal(t, poll.CreatedBy, result.CreatedBy)
	require.Equal(t, poll.CreatedAt, result.CreatedAt)

}

func TestUpdatePollStatus(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	args := UpdatePollStatusParams{
		ID:       poll.ID,
		IsActive: false,
	}

	result, err := testStore.UpdatePollStatus(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, poll.ID, result.ID)
	require.Equal(t, user.ID, poll.CreatedBy)
	require.Equal(t, poll.CreatedAt, poll.CreatedAt)
	require.NotEqual(t, poll.IsActive, result.IsActive)
	require.Equal(t, poll.CreatedBy, result.CreatedBy)
}

func TestGetPolls(t *testing.T) {
	user := createRandomUser(t)
	poll_arr := make([]CreatePollRow, 0)
	for i := 0; i < 10; i++ {
		poll := createRandomPoll(t, user.ID)
		poll_arr = append(poll_arr, poll)
	}

	polls, err := testStore.GetPolls(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, polls)
	require.GreaterOrEqual(t, len(polls), len(poll_arr))
}
