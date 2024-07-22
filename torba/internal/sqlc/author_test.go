package sqlc

import (
	"context"
	"openSo/internal/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func CreateRandomAuthor(t *testing.T) Author {
	arg := CreateFullAuthorParams{
		NickName:        util.RandomString(8),
		Email:           util.RandomString(8) + "@gmail.com",
		PasswordHash:    util.RandomString(8),
		Payments:        util.RandomString(8),
		Bio:             util.RandomString(12),
		Link:            util.RandomString(5),
		ProfileImageUrl: util.RandomString(5),
		AdditionalInfo:  util.RandomString(3),
	}

	author, err := test.CreateFullAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.NickName, author.NickName)
	require.Equal(t, arg.Email, author.Email)
	require.Equal(t, arg.PasswordHash, author.PasswordHash)
	require.Equal(t, arg.Payments, author.Payments)
	require.Equal(t, arg.Bio, author.Bio)
	require.Equal(t, arg.Link, author.Link)
	require.Equal(t, arg.ProfileImageUrl, author.ProfileImageUrl)
	require.Equal(t, arg.AdditionalInfo, author.AdditionalInfo)

	return author
}

func TestSignAuthor(t *testing.T) {
	arg := CreateAuthorParams{
		NickName:     util.RandomString(12),
		Email:        util.RandomString(12) + "@gmail.com",
		PasswordHash: util.RandomString(12),
	}

	author, err := test.CreateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, author)

	require.Equal(t, arg.NickName, author.NickName)
	require.NotEmpty(t, author.AuthorID)
}

func TestSignFullAuthor(t *testing.T) {
	CreateRandomAuthor(t)
}

func TestDeleteAuthor(t *testing.T) {
	author := CreateRandomAuthor(t)

	err := test.DeleteAuthor(context.Background(), author.AuthorID)
	require.NoError(t, err)
}

func TestGetAuthorByAuthorID(t *testing.T) {
	author := CreateRandomAuthor(t)

	AuthorID, err := test.GetAuthorByID(context.Background(), author.AuthorID)
	require.NoError(t, err)

	require.NotEmpty(t, AuthorID)
	require.Equal(t, author.NickName, AuthorID.NickName)
	require.Equal(t, author.Email, AuthorID.Email)
	require.Equal(t, author.PasswordHash, AuthorID.PasswordHash)
	require.Equal(t, author.Payments, AuthorID.Payments)
	require.Equal(t, author.Bio, AuthorID.Bio)
	require.Equal(t, author.Link, AuthorID.Link)
	require.Equal(t, author.ProfileImageUrl, AuthorID.ProfileImageUrl)
	require.Equal(t, author.AdditionalInfo, AuthorID.AdditionalInfo)

}

func TestGetAuthorByName(t *testing.T) {
	author := CreateRandomAuthor(t)

	AuthorID, err := test.GetAuthorByName(context.Background(), author.NickName)
	require.NoError(t, err)

	require.NotEmpty(t, AuthorID)
	require.Equal(t, author.NickName, AuthorID.NickName)
	require.Equal(t, author.Email, AuthorID.Email)
	require.Equal(t, author.PasswordHash, AuthorID.PasswordHash)
	require.Equal(t, author.Payments, AuthorID.Payments)
	require.Equal(t, author.Bio, AuthorID.Bio)
	require.Equal(t, author.Link, AuthorID.Link)
	require.Equal(t, author.ProfileImageUrl, AuthorID.ProfileImageUrl)
	require.Equal(t, author.AdditionalInfo, AuthorID.AdditionalInfo)

}

func TestListAuthorAuthorID(t *testing.T) {
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

func TestUpdateAuthor(t *testing.T) {
	author := CreateRandomAuthor(t)

	arg := UpdateAuthorParams{
		AuthorID:     author.AuthorID,
		NickName:     util.RandomString(8),
		Email:        util.RandomString(8) + "@gmail.com",
		PasswordHash: util.RandomString(8),
	}

	a, err := test.UpdateAuthor(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, a.AuthorID)
	require.NotEmpty(t, a.NickName)

	require.NotEqual(t, author.NickName, a.NickName)
}

func TestUpdateAuthorFull(t *testing.T) {
	author := CreateRandomAuthor(t)

	arg := UpdateAuthorFullParams{
		AuthorID:        author.AuthorID,
		NickName:        util.RandomString(7),
		Email:           util.RandomString(6) + "@gmail.com",
		PasswordHash:    util.RandomString(10),
		Payments:        util.RandomString(12),
		Bio:             util.RandomString(6),
		Link:            util.RandomString(4),
		ProfileImageUrl: util.RandomString(6),
	}

	a, err := test.UpdateAuthorFull(context.Background(), arg)
	require.NoError(t, err)

	require.NotEqual(t, author.NickName, a.NickName)
	require.Equal(t, a.AuthorID, a.AuthorID)
}
