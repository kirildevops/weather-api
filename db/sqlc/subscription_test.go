package db

import (
	"context"
	"testing"

	"github.com/kirildevops/weather-api/util"
	"github.com/stretchr/testify/require"
)

func createRandomSubscription(t *testing.T) Subscription {
	arg := InsertSubscriptionParams{
		Email:     util.RandomEmail(),
		City:      util.RandomCity(),
		Frequency: FrequencyEnum(util.RandomFrequency()),
	}

	sub, err := testQueries.InsertSubscription(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sub)

	require.Equal(t, arg.Email, sub.Email)
	require.Equal(t, arg.City, sub.City)
	require.Equal(t, arg.Frequency, sub.Frequency)

	require.NotZero(t, sub.ID)
	require.NotZero(t, sub.Token)

	return sub
}

func TestInsertSubscription(t *testing.T) {
	createRandomSubscription(t)
}

func TestGetSubscription(t *testing.T) {
	sub1 := createRandomSubscription(t)
	sub2, err := testQueries.GetSubscription(context.Background(), sub1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, sub2)

	require.Equal(t, sub1.ID, sub2.ID)
	require.Equal(t, sub1.City, sub2.City)
	require.Equal(t, sub1.Token, sub2.Token)
	require.Equal(t, sub1.Confirmed, sub2.Confirmed)
}
