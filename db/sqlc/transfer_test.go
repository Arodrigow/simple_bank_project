package db

import (
	"context"
	"testing"
	"time"

	"github.com/Arodrigow/simple_bank_project/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, toAccount, fromAccount Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomMoneyValue(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)
	createRandomTransfer(t, toAccount, fromAccount)
}

func TestGetTransfer(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)
	mockTransfer := createRandomTransfer(t, toAccount, fromAccount)

	resultTransfer, err := testQueries.GetTransfer(context.Background(), mockTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, resultTransfer)

	require.Equal(t, mockTransfer.ID, resultTransfer.ID)
	require.Equal(t, mockTransfer.ToAccountID, resultTransfer.ToAccountID)
	require.Equal(t, mockTransfer.FromAccountID, resultTransfer.FromAccountID)
	require.Equal(t, mockTransfer.Amount, resultTransfer.Amount)
	require.WithinDuration(t, mockTransfer.CreatedAt, resultTransfer.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	toAccount := createRandomAccount(t)
	fromAccount := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, toAccount, fromAccount)
		createRandomTransfer(t, fromAccount, toAccount)
	}

	arg := ListTransfersParams{
		FromAccountID: toAccount.ID,
		ToAccountID:   toAccount.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
		require.True(t, transfer.ToAccountID == fromAccount.ID || transfer.FromAccountID == fromAccount.ID)
	}
}
