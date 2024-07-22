package sqlc

import (
	"context"
	"openSo/internal/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:        util.UserName(),
		Email:           util.Email(),
		PasswordHash:    util.Password(),
		ProfileImageUrl: util.RandomString(7),
	}

	user, err := test.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)

	require.NotZero(t, user.UserID)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByUserID(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := test.GetUserByID(context.Background(), user1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.ProfileImageUrl, user2.ProfileImageUrl)
}

func TestGetUserByName(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := test.GetUserByName(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.ProfileImageUrl, user2.ProfileImageUrl)
}

func TestListUserName(t *testing.T) {
	var users []User

	users, err := test.ListUserName(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.NotEmpty(t, user.Username)
		require.NotEmpty(t, user.CreatedAt)
		require.NotEmpty(t, user.Email)
		require.NotEmpty(t, user.PasswordHash)
	}
}

func TestListUserUserID(t *testing.T) {
	var users []User

	users, err := test.ListUserID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, users)

	for _, user := range users {
		require.NotEmpty(t, user)
		require.NotEmpty(t, user.Username)
		require.NotEmpty(t, user.CreatedAt)
		require.NotEmpty(t, user.Email)
		require.NotEmpty(t, user.PasswordHash)
	}
}

func TestUpdateUser(t *testing.T) {
	account1 := CreateRandomUser(t)

	arg := UpdateUserParams{
		UserID:       account1.UserID,
		Username:     util.UserName(),
		Email:        util.Email(),
		PasswordHash: util.Password(),
	}

	usr, err := test.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, usr.Username)
	require.NotEmpty(t, usr.UserID)

	require.Equal(t, account1.UserID, usr.UserID)
}

func TestUpdateUserFull(t *testing.T) {
	account1 := CreateRandomUser(t)

	arg := UpdateUserParams{
		UserID:          account1.UserID,
		Username:        util.UserName(),
		Email:           util.Email(),
		PasswordHash:    util.Password(),
		ProfileImageUrl: util.RandomString(7),
	}

	usr, err := test.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, usr.Username)
	require.NotEmpty(t, usr.UserID)

	require.Equal(t, account1.UserID, usr.UserID)
}

func TestDeleteUser(t *testing.T) {
	account1 := CreateRandomUser(t)

	err := test.DeleteUser(context.Background(), account1.UserID)
	require.NoError(t, err)

	u, err := test.GetUserByID(context.Background(), account1.UserID)
	require.Error(t, err)
	require.Empty(t, u)
}
