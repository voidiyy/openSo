package sqlc

import (
	"context"
	"openSo/internal/util"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateProject(t *testing.T) Project {
	var err error

	p := Project{}

	str := "1234.13"
	var fund pgtype.Numeric

	err = fund.Scan(str)
	if err != nil {
		t.Fatal(err)
	}

	a := CreateRandomAuthor(t)

	arg := CreateProjectParams{
		AuthorID:    a.AuthorID,
		Title:       util.RandomString(10),
		Category:    util.RandomString(10),
		Description: util.RandomString(10),
		Link:        util.RandomString(8),
		Details:     util.RandomString(8),
		Payments:    util.RandomString(6),
		FundingGoal: fund,
	}

	p, err = test.CreateProject(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, p)
	require.Equal(t, a.AuthorID, p.AuthorID)

	return p
}

func TestProjects_CreateProject(t *testing.T) {
	CreateProject(t)
}

func TestDeleteProject(t *testing.T) {

	d := CreateProject(t)
	err := test.DeleteProject(context.Background(), d.ProjectID)
	require.NoError(t, err)
}

func TestProjects_GetProjectByAuthor(t *testing.T) {
	p := CreateProject(t)

	get, err := test.GetProjectByAuthor(context.Background(), p.AuthorID)
	require.NoError(t, err)
	require.NotEmpty(t, get)

	require.Equal(t, p.ProjectID, get.ProjectID)
	require.Equal(t, p.AuthorID, get.AuthorID)
	require.Equal(t, p.Title, get.Title)
	require.Equal(t, p.Description, get.Description)
	require.Equal(t, p.Link, get.Link)
	require.Equal(t, p.Details, get.Details)
	require.Equal(t, p.Payments, get.Payments)
	require.Equal(t, p.FundingGoal, get.FundingGoal)
}

func TestProjects_GetProjectByID(t *testing.T) {
	p := CreateProject(t)

	get, err := test.GetProjectByCategory(context.Background(), p.Category)
	require.NoError(t, err)
	require.NotEmpty(t, get)

	require.Equal(t, p.ProjectID, get.ProjectID)
	require.Equal(t, p.AuthorID, get.AuthorID)
	require.Equal(t, p.Title, get.Title)
	require.Equal(t, p.Description, get.Description)
	require.Equal(t, p.Link, get.Link)
	require.Equal(t, p.Details, get.Details)
	require.Equal(t, p.Payments, get.Payments)
	require.Equal(t, p.FundingGoal, get.FundingGoal)
}

func TestProjects_GetProjectByTitle(t *testing.T) {
	p := CreateProject(t)

	get, err := test.GetProjectByTitle(context.Background(), p.Title)
	require.NoError(t, err)
	require.NotEmpty(t, get)

	require.Equal(t, p.ProjectID, get.ProjectID)
	require.Equal(t, p.AuthorID, get.AuthorID)
	require.Equal(t, p.Title, get.Title)
	require.Equal(t, p.Description, get.Description)
	require.Equal(t, p.Link, get.Link)
	require.Equal(t, p.Details, get.Details)
	require.Equal(t, p.Payments, get.Payments)
	require.Equal(t, p.FundingGoal, get.FundingGoal)
}

func TestProjects_ListPrjAuthor(t *testing.T) {
	var projects []Project
	var err error

	projects, err = test.ListProjectAuthor(context.Background())
	require.NoError(t, err)

	require.NotEmpty(t, projects)

	for _, project := range projects {
		require.NotEmpty(t, project)
	}
}
