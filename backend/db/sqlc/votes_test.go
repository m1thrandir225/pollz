package db

import (
	"context"
	"m1thrandir225/cicd2025/util"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomVote(t *testing.T, optionId uuid.UUID) Vote {
	args := CreateVoteParams{
		IpAddress: util.RandomString(11),
		UserAgent: util.RandomString(32),
		OptionID:  optionId,
	}

	vote, err := testStore.CreateVote(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, vote)

	require.Equal(t, vote.IpAddress, args.IpAddress)
	require.Equal(t, vote.OptionID, args.OptionID)
	require.Equal(t, vote.UserAgent, args.UserAgent)

	require.NotEmpty(t, vote.ID)
	require.NotEmpty(t, vote.VotedAt)

	return vote
}

func TestCreateVote(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)
	createRandomVote(t, option.ID)
}

func TestGetVote(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)
	vote := createRandomVote(t, option.ID)

	result, err := testStore.GetVote(context.Background(), vote.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, vote.ID, result.ID)
	require.Equal(t, vote.IpAddress, result.IpAddress)
	require.Equal(t, vote.OptionID, result.OptionID)
	require.WithinDuration(t, vote.VotedAt, result.VotedAt, time.Millisecond)
}

func TestUpdateVote(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)
	vote := createRandomVote(t, option.ID)

	newOption := createRandomPollOption(t, poll.ID)
	arg := UpdateVoteOptionParams{
		ID:       vote.ID,
		OptionID: newOption.ID,
	}

	updated, err := testStore.UpdateVoteOption(t.Context(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updated)

	require.Equal(t, vote.ID, updated.ID)
	require.Equal(t, vote.IpAddress, updated.IpAddress)
	require.NotEqual(t, vote.OptionID, updated.OptionID)
	require.WithinDuration(t, vote.VotedAt, updated.VotedAt, time.Millisecond)

	require.Equal(t, updated.OptionID, arg.OptionID)
}

func TestDeleteVote(t *testing.T) {
	user := createRandomUser(t)
	poll := createRandomPoll(t, user.ID)
	option := createRandomPollOption(t, poll.ID)
	vote := createRandomVote(t, option.ID)

	deleted, err := testStore.DeleteVote(context.Background(), vote.ID)
	require.NoError(t, err)
	require.NotEmpty(t, deleted)

	require.Equal(t, vote.ID, deleted.ID)
	require.Equal(t, vote.IpAddress, deleted.IpAddress)
	require.Equal(t, vote.OptionID, deleted.OptionID)
	require.WithinDuration(t, vote.VotedAt, deleted.VotedAt, time.Millisecond)

	result, err := testStore.GetVote(context.Background(), deleted.ID)
	require.Error(t, err)
	require.Empty(t, result)
}
