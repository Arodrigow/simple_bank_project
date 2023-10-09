package db

import (
	"context"
	"testing"
	"time"

	"github.com/Arodrigow/simple_bank_project/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "secret",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	mockUser := createRandomUser(t)
	resultUser, err := testQueries.GetUser(context.Background(), mockUser.Username)
	require.NoError(t, err)
	require.NotEmpty(t, resultUser)

	require.Equal(t, mockUser.Username, resultUser.Username)
	require.Equal(t, mockUser.HashedPassword, resultUser.HashedPassword)
	require.Equal(t, mockUser.FullName, resultUser.FullName)
	require.Equal(t, mockUser.Email, resultUser.Email)

	require.WithinDuration(t, mockUser.PasswordChangedAt, resultUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, mockUser.CreatedAt, resultUser.CreatedAt, time.Second)
}
