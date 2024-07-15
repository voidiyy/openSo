package sqlc

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"torba/internal/util"
)

func CreateRandomAuthor(t *testing.T) Author {
	arg := SignFullAuthorParams{
		NickName:        util.RandomString(8),
		Email:           util.RandomString(8) + "@gmail.com",
		PasswordHash:    util.RandomString(8),
		Payments:        util.RandomString(8),
		Projects:        util.RandomStringSlice(4),
		Bio:             util.RandomString(12),
		Link:            util.RandomString(5),
		ProfileImageUrl: util.RandomString(5),
		AdditionalInfo:  util.RandomStringSlice(3),
	}

	author, err := test.SignFullAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.NickName, author.NickName)
	require.Equal(t, arg.Email, author.Email)
	require.Equal(t, arg.PasswordHash, author.PasswordHash)
	require.Equal(t, arg.Payments, author.Payments)
	require.ElementsMatch(t, arg.Projects, author.Projects)
	require.Equal(t, arg.Bio, author.Bio)
	require.Equal(t, arg.Link, author.Link)
	require.Equal(t, arg.ProfileImageUrl, author.ProfileImageUrl)
	require.ElementsMatch(t, arg.AdditionalInfo, author.AdditionalInfo)

	return author
}

func TestSignAuthor(t *testing.T) {
	arg := SignAuthorParams{
		NickName:     util.RandomString(12),
		Email:        util.RandomString(12),
		PasswordHash: util.RandomString(12),
	}

	author, err := test.SignAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.NickName, author.NickName)
	require.NotEmpty(t, author.ID)
}

func TestSignFullAuthor(t *testing.T) {
	CreateRandomAuthor(t)
}

func TestDeleteAuthor(t *testing.T) {
	author := CreateRandomAuthor(t)

	err := test.DeleteAuthor(context.Background(), author.ID)
	require.NoError(t, err)
}

func TestGetAuthorByID(t *testing.T) {
	author := CreateRandomAuthor(t)

	id, err := test.GetAuthorByID(context.Background(), author.ID)
	require.NoError(t, err)

	require.NotEmpty(t, id)
	require.Equal(t, author.NickName, id.NickName)
	require.Equal(t, author.Email, id.Email)
	require.Equal(t, author.PasswordHash, id.PasswordHash)
	require.Equal(t, author.Payments, id.Payments)
	require.ElementsMatch(t, author.Projects, id.Projects)
	require.Equal(t, author.Bio, id.Bio)
	require.Equal(t, author.Link, id.Link)
	require.Equal(t, author.ProfileImageUrl, id.ProfileImageUrl)
	require.ElementsMatch(t, author.AdditionalInfo, id.AdditionalInfo)

}

func TestGetAuthorByName(t *testing.T) {
	author := CreateRandomAuthor(t)

	id, err := test.GetAuthorByName(context.Background(), author.NickName)
	require.NoError(t, err)

	require.NotEmpty(t, id)
	require.Equal(t, author.NickName, id.NickName)
	require.Equal(t, author.Email, id.Email)
	require.Equal(t, author.PasswordHash, id.PasswordHash)
	require.Equal(t, author.Payments, id.Payments)
	require.ElementsMatch(t, author.Projects, id.Projects)
	require.Equal(t, author.Bio, id.Bio)
	require.Equal(t, author.Link, id.Link)
	require.Equal(t, author.ProfileImageUrl, id.ProfileImageUrl)
	require.ElementsMatch(t, author.AdditionalInfo, id.AdditionalInfo)

}

func TestListAuthorID(t *testing.T) {
	var authors []Author

	authors, err := test.ListAuthorID(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, authors)

	for _, a := range authors {
		require.NotEmpty(t, a.NickName)
		require.NotEmpty(t, a.Email)
		require.NotEmpty(t, a.PasswordHash)
	}
}

func TestListAuthorName(t *testing.T) {
	var authors []Author

	authors, err := test.ListAuthorName(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, authors)

	for _, a := range authors {
		require.NotEmpty(t, a.NickName)
		require.NotEmpty(t, a.Email)
		require.NotEmpty(t, a.PasswordHash)
	}
}
