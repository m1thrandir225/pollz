package db

import (
	"context"
	"m1thrandir225/cicd2025/util"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
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

func TestCreateMultipleOptions(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	options := []string{"option 1", "option 2", "option 3"}

	arg := CreateMultipleOptionsParams{
		PollID:      poll.ID,
		OptionTexts: options,
	}

	result, err := testStore.CreateMultipleOptions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, len(options), len(result))
	for _, v := range result {
		require.Equal(t, v.PollID, poll.ID)
	}
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

func TestGetOptionsForPoll(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)

	options := make([]PollOption, 10)
	for i := range options {
		option := createRandomPollOption(t, poll.ID)
		options[i] = option
	}

	result, err := testStore.GetOptionsForPoll(context.Background(), poll.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, len(options), len(result))

	for _, v := range result {
		require.Equal(t, v.PollID, poll.ID)
	}
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
