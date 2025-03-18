package db

import (
	"context"
	"m1thrandir225/cicd2025/util"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomPoll(t *testing.T, userId uuid.UUID) Poll {
	pollDescription := util.RandomString(120)

	activeTime := time.Now()

	args := CreatePollParams{
		Description: pollDescription,
		CreatedBy:   userId,
		ActiveUntil: activeTime,
	}
	poll, err := testStore.CreatePoll(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, poll)
	require.Equal(t, pollDescription, poll.Description)
	require.Equal(t, userId, poll.CreatedBy)
	require.WithinDuration(t, poll.ActiveUntil, activeTime, time.Second)
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
	require.WithinDuration(t, poll.ActiveUntil, result.ActiveUntil, time.Microsecond)
}

func TestUpdatePoll(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	newDescription := util.RandomString(120)

	newActiveTime := time.Now().Add(time.Hour * 24 * 7)
	updateArgs := UpdatePollParams{
		ID:          poll.ID,
		Description: newDescription,
		ActiveUntil: newActiveTime,
	}
	result, err := testStore.UpdatePoll(context.Background(), updateArgs)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, poll.ID, result.ID)
	require.Equal(t, newDescription, result.Description)
	require.NotEqual(t, poll.Description, result.Description)
	require.WithinDuration(t, result.ActiveUntil, newActiveTime, time.Second)
	require.NotEqual(t, poll.ActiveUntil, result.ActiveUntil)
	require.Equal(t, poll.CreatedBy, result.CreatedBy)
	require.Equal(t, poll.CreatedAt, result.CreatedAt)
}

func TestUpdatePollStatus(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	newActiveTime := time.Now()
	args := UpdatePollStatusParams{
		ID:          poll.ID,
		ActiveUntil: newActiveTime,
	}

	result, err := testStore.UpdatePollStatus(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, poll.ID, result.ID)
	require.Equal(t, user.ID, poll.CreatedBy)
	require.Equal(t, poll.CreatedAt, poll.CreatedAt)
	require.NotEqual(t, poll.ActiveUntil, result.ActiveUntil)
	require.WithinDuration(t, result.ActiveUntil, newActiveTime, time.Microsecond)
	require.Equal(t, poll.CreatedBy, result.CreatedBy)
}

func TestGetPolls(t *testing.T) {
	user := createRandomUser(t)
	poll_arr := make([]Poll, 0)
	for i := 0; i < 10; i++ {
		poll := createRandomPoll(t, user.ID)
		poll_arr = append(poll_arr, poll)
	}

	polls, err := testStore.GetPolls(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, polls)
	require.Equal(t, len(poll_arr), len(polls))
}
