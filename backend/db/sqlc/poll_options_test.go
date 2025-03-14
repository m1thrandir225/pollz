package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"m1thrandir225/cicd2025/util"
	"testing"
)

func createRandomPollOption(t *testing.T, pollId uuid.UUID) PollOption {
	optionText := util.RandomString(10)
	args := CreatePollOptionParams{
		PollID:     pollId,
		OptionText: optionText,
	}

	result, err := testStore.CreatePollOption(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, pollId, result.PollID)
	require.Equal(t, optionText, result.OptionText)

	return result
}

func TestCreatePollOption(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	createRandomPollOption(t, poll.ID)
}

func TestGetOption(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)

	result, err := testStore.GetOption(context.Background(), option.ID)
	require.NoError(t, err)
	require.Equal(t, option.ID, result.ID)
	require.Equal(t, option.OptionText, result.OptionText)
	require.Equal(t, option.CreatedAt, result.CreatedAt)
	require.Equal(t, option.PollID, result.PollID)
	require.Equal(t, poll.ID, result.PollID)
}

func TestUpdatePollOption(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)

	updateOptionText := util.RandomString(10)

	args := UpdatePollOptionParams{
		ID:         option.ID,
		OptionText: updateOptionText,
	}

	result, err := testStore.UpdatePollOption(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, option.ID, result.ID)
	require.NotEqual(t, option.OptionText, result.OptionText)
	require.Equal(t, updateOptionText, result.OptionText)
	require.Equal(t, option.CreatedAt, result.CreatedAt)
	require.Equal(t, option.PollID, result.PollID)
	require.Equal(t, poll.ID, result.PollID)
}

func TestDeletePollOption(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)

	res, err := testStore.DeletePollOption(context.Background(), option.ID)
	require.NoError(t, err)
	require.NotEmpty(t, res)
	require.Equal(t, option.ID, res.ID)
	require.Equal(t, poll.ID, res.PollID)
	require.Equal(t, option.CreatedAt, res.CreatedAt)
	require.Equal(t, option.PollID, res.PollID)
	require.Equal(t, option.OptionText, res.OptionText)

	deleted, err := testStore.GetOption(context.Background(), option.ID)
	require.Error(t, err)
	require.Empty(t, deleted)
}
