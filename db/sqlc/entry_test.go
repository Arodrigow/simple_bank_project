package db

import (
	"context"
	"testing"
	"time"

	"github.com/Arodrigow/simple_bank_project/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, mockAccount Account) Entry {
	arg := CreateEntryParams{
		AccountID: mockAccount.ID,
		Amount:    util.RandomMoneyValue(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, mockAccount.ID, entry.AccountID)
	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	mockAccount := createRandomAccount(t)
	createRandomEntry(t, mockAccount)
}

func TestGetEntry(t *testing.T) {
	mockAccount := createRandomAccount(t)
	mockEntry := createRandomEntry(t, mockAccount)
	resultEntry, err := testQueries.GetEntry(context.Background(), mockEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, resultEntry)

	require.Equal(t, mockEntry.ID, resultEntry.ID)
	require.Equal(t, mockEntry.AccountID, resultEntry.AccountID)
	require.Equal(t, mockEntry.Amount, resultEntry.Amount)
	require.WithinDuration(t, mockEntry.CreatedAt, resultEntry.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	mockAccount := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomEntry(t, mockAccount)
	}

	args := ListEntriesParams{
		AccountID: mockAccount.ID,
		Offset:    5,
		Limit:     5,
	}

	entries, err := testQueries.ListEntries(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, args.AccountID, entry.AccountID)
	}
}
