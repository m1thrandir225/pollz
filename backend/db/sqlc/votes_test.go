package db

import (
	"context"
	"m1thrandir225/cicd2025/util"
	"testing"

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

func TestGetVote(t *testing.T) {}

func TestUpdateVote(t *testing.T) {}

func TestDeleteVote(t *testing.T) {}
