package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"torba/internal/util"
)

func CreateRandomUser(t *testing.T) User {
	arg := SignFullUserParams{
		Username:          util.UserName(),
		Email:             util.Email(),
		PasswordHash:      util.Password(),
		DonationSum:       util.RandomFloat(3, 12),
		SupportedProjects: util.RandomIntSlice(2, 3, 16),
		ProfileImageUrl:   util.RandomString(7),
	}

	user, err := test.SignFullUser(context.Background(), arg) // змініть на реальний виклик вашої функції створення користувача
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)
	require.Equal(t, arg.DonationSum, user.DonationSum)
	require.ElementsMatch(t, arg.SupportedProjects, user.SupportedProjects)
	require.Equal(t, arg.ProfileImageUrl, user.ProfileImageUrl)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestSignUser(t *testing.T) {
	arg := SignUserParams{
		Username:     util.UserName(),
		Email:        util.Email(),
		PasswordHash: util.Password(),
	}

	user, err := test.SignUser(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.Username, user.Username)
}

func TestSignFullUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByID(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := test.GetUserByID(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.DonationSum, user2.DonationSum)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.ProfileImageUrl, user2.ProfileImageUrl)
	require.Equal(t, user1.SupportedProjects, user2.SupportedProjects)

}

func TestGetUserByName(t *testing.T) {
	user1 := CreateRandomUser(t)
	user2, err := test.GetUserByName(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.CreatedAt, user2.CreatedAt)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.Equal(t, user1.ProfileImageUrl, user2.ProfileImageUrl)
	require.Equal(t, user1.DonationSum, user2.DonationSum)
	require.Equal(t, user1.SupportedProjects, user2.SupportedProjects)
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

func TestListUserID(t *testing.T) {
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
		ID:           account1.ID,
		Username:     util.UserName(),
		Email:        util.Email(),
		PasswordHash: util.Password(),
	}

	usr, err := test.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, usr.Username)
	require.NotEmpty(t, usr.ID)

	require.Equal(t, account1.ID, usr.ID)
}

func TestUpdateUserFull(t *testing.T) {
	account1 := CreateRandomUser(t)

	arg := UpdateFullParams{
		ID:                account1.ID,
		Username:          util.UserName(),
		Email:             util.Email(),
		PasswordHash:      util.Password(),
		DonationSum:       util.RandomFloat(3, 10),
		SupportedProjects: util.RandomIntSlice(3, 5, 10),
		ProfileImageUrl:   util.RandomString(7),
	}

	usr, err := test.UpdateFull(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, usr.Username)
	require.NotEmpty(t, usr.ID)

	require.Equal(t, account1.ID, usr.ID)
}

func TestDeleteUser(t *testing.T) {
	account1 := CreateRandomUser(t)

	err := test.DeleteUser(context.Background(), account1.ID)
	require.NoError(t, err)

	u, err := test.GetUserByID(context.Background(), account1.ID)
	require.Error(t, err)
	require.Empty(t, u)
}
