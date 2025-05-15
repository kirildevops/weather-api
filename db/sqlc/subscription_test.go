package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertSubscription(t *testing.T) {
	arg := InsertSubscriptionParams{
		Email:     "test@test.abc",
		City:      "Neverland",
		Frequency: "daily",
	}

	sub, err := testQueries.InsertSubscription(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, sub)

	require.Equal(t, arg.Email, sub.Email)
	require.Equal(t, arg.City, sub.City)
	require.Equal(t, arg.Frequency, sub.Frequency)

	require.NotZero(t, sub.ID)
	require.NotZero(t, sub.Token)
}
