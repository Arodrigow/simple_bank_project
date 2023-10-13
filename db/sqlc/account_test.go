package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Arodrigow/simple_bank_project/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoneyValue(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	mockAccount := createRandomAccount(t)
	resultAccount, err := testQueries.GetAccount(context.Background(), mockAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, resultAccount)

	require.Equal(t, mockAccount.ID, resultAccount.ID)
	require.Equal(t, mockAccount.Owner, resultAccount.Owner)
	require.Equal(t, mockAccount.Balance, resultAccount.Balance)
	require.Equal(t, mockAccount.Currency, resultAccount.Currency)
	require.WithinDuration(t, mockAccount.CreatedAt, resultAccount.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	mockAccount := createRandomAccount(t)

	args := UpdateAccountParams{
		ID:      mockAccount.ID,
		Balance: util.RandomMoneyValue(),
	}

	resultAccount, err := testQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, resultAccount)

	require.Equal(t, mockAccount.ID, resultAccount.ID)
	require.Equal(t, mockAccount.Owner, resultAccount.Owner)
	require.Equal(t, args.Balance, resultAccount.Balance)
	require.Equal(t, mockAccount.Currency, resultAccount.Currency)
	require.WithinDuration(t, mockAccount.CreatedAt, resultAccount.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	mockAccount := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), mockAccount.ID)
	require.NoError(t, err)

	mockAccount2, err := testQueries.GetAccount(context.Background(), mockAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, mockAccount2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	args := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Offset: 0,
		Limit:  5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}
