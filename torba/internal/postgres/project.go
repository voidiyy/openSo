package postgres

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgtype"
)

func (p *DB) CategoryExistsPrj(ctx context.Context, category string) error {
	query := "SELECT EXISTS (SELECT 1 FROM projects WHERE category = $1)"
	var exists bool
	err := p.pdb.QueryRow(ctx, query, category).Scan(&exists)
	if err != nil {
		return errors.New("error check category existence: " + err.Error())
	}
	if !exists {
		return errors.New("category does not exist")
	}

	return nil
}

func (p *DB) CategoryAndSubCategoryExistsPrj(ctx context.Context, category, subcategory string) error {
	query := "SELECT EXISTS (SELECT 1 FROM projects WHERE category = $1 and sub_category = $2)"

	var exists bool
	err := p.pdb.QueryRow(ctx, query, category, subcategory).Scan(&exists)
	if err != nil {
		return errors.New("error check sub_category existence: " + err.Error())
	}
	if !exists {
		return errors.New("sub_category does not exist")
	}

	return nil
}

type CreateProjectParams struct {
	AuthorID    int64
	Title       string
	Category    string
	Subcategory string
	Description string
	Link        string
	Details     string
	Payments    string
	FundingGoal string
}

func (p *DB) CreateProject(ctx context.Context, arg *CreateProjectParams) error {
	query := "insert into projects (author_id, title, category, sub_category, description, link,details, payments, funding_goal, created_at) values ($1, $2,$3,$4,$5,$6,$7,$8,now())"

	_, err := p.pdb.Exec(ctx, query, arg.AuthorID, arg.Title, arg.Category, arg.Subcategory, arg.Description, arg.Link, arg.Details, arg.Payments, arg.FundingGoal)
	if err != nil {
		p.errl.Printf("error creating query: %v", err)
		return err
	}
	p.infol.Printf("created project: %v", arg.Title)
	return nil
}

func (p *DB) DeleteProject(ctx context.Context, id int64) error {
	query := "delete from projects where project_id = $1"
	_, err := p.pdb.Exec(ctx, query, id)
	if err != nil {
		p.errl.Printf("error deleting project %d: %v", id, err)
		return err
	}
	p.infol.Printf("deleted project: %v", id)
	return nil
}

type UpdateProjectParams struct {
	ProjectID   int32
	Title       string
	Category    string
	Subcategory string
	Description string
	Link        string
	Details     string
	Payments    string
	Status      bool
	FundingGoal string
}

func (p *DB) UpdateProject(ctx context.Context, arg *UpdateProjectParams) error {
	query := "update projects set title = $2, category = $3, sub_category= $4, description = $5, link = $6, details = $7, payments = $8, status = $9, funding_goal = $10, updated_at = now() where project_id = $1"

	_, err := p.pdb.Exec(ctx, query, arg.ProjectID, arg.Title, arg.Category, arg.Subcategory, arg.Description, arg.Link, arg.Details, arg.Payments, arg.Status, arg.FundingGoal)
	if err != nil {
		p.errl.Printf("error updating project %d: %v", arg.ProjectID, err)
		return err
	}
	p.infol.Printf("updated project: %v", arg.Title)
	return nil
}

func (p *DB) ListProjectsByCategory(ctx context.Context, category string) ([]Project, error) {
	query := "SELECT * FROM projects WHERE category = $1"

	rows, err := p.pdb.Query(ctx, query, category)
	if err != nil {
		p.errl.Printf("failed to query projects by category: %v", err)
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err := rows.Scan(
			&project.AuthorID,
			&project.Title,
			&project.Category,
			&project.Subcategory,
			&project.Description,
			&project.Link,
			&project.Details,
			&project.Payments,
			&project.FundingGoal,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			p.errl.Printf("failed to scan project: %v", err)
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		p.errl.Printf("error occurred during iteration: %v", err)
		return nil, err
	}

	p.infol.Printf("projects listed by category: %v", category)
	return projects, nil
}

func (p *DB) ListProjectsByCategoryAndSubcategory(ctx context.Context, category, subcategory string) ([]Project, error) {
	query := "SELECT * FROM projects WHERE category = $1 AND sub_category = $2"

	rows, err := p.pdb.Query(ctx, query, category, subcategory)
	if err != nil {
		p.errl.Printf("failed to query projects by category and subcategory: %v", err)
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err := rows.Scan(
			&project.AuthorID,
			&project.Title,
			&project.Category,
			&project.Subcategory,
			&project.Description,
			&project.Link,
			&project.Details,
			&project.Payments,
			&project.FundingGoal,
			&project.CreatedAt,
			&project.UpdatedAt,
		)
		if err != nil {
			p.errl.Printf("failed to scan project: %v", err)
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		p.errl.Printf("error occurred during iteration: %v", err)
		return nil, err
	}

	p.infol.Printf("projects listed by category: %v and subcategory: %v", category, subcategory)
	return projects, nil
}

func (p *DB) GetProjectByID(ctx context.Context, id int32) (Project, error) {
	query := "select * from projects where project_id = $1"

	row := p.pdb.QueryRow(ctx, query, id)
	var project Project
	err := row.Scan(
		&project.AuthorID,
		&project.AuthorID,
		&project.Title,
		&project.Category,
		&project.Subcategory,
		&project.Description,
		&project.Link,
		&project.Details,
		&project.Payments,
		&project.FundingGoal,
		&project.CreatedAt,
		&project.UpdatedAt)
	if err != nil {
		return Project{}, err
	}
	p.infol.Printf("goted project by ID: %v", project.Title)
	return project, nil
}

func (p *DB) GetProjectByAuthor(ctx context.Context, id int64) (Project, error) {
	query := "select * from projects where author_id = $1"
	row := p.pdb.QueryRow(ctx, query, id)
	var project Project
	err := row.Scan(
		&project.ProjectID,
		&project.AuthorID,
		&project.Title,
		&project.Category,
		&project.Subcategory,
		&project.Description,
		&project.Link,
		&project.Details,
		&project.Payments,
		&project.FundingGoal,
		&project.FundsRaised,
		&project.CreatedAt,
		&project.UpdatedAt)
	if err != nil {
		return Project{}, err
	}

	p.infol.Printf("goted project by author_ID: %v", project.AuthorID)
	return project, nil
}

func (p *DB) GetProjectByTitle(ctx context.Context, title string) (Project, error) {
	query := "select * from projects where title = $1"
	row := p.pdb.QueryRow(ctx, query, title)
	var project Project
	err := row.Scan(
		&project.ProjectID,
		&project.AuthorID,
		&project.Title,
		&project.Category,
		&project.Subcategory,
		&project.Description,
		&project.Link,
		&project.Details,
		&project.Payments,
		&project.FundingGoal,
		&project.FundsRaised,
		&project.CreatedAt,
		&project.UpdatedAt)
	if err != nil {
		return Project{}, err
	}

	p.infol.Printf("goted project by title: %v", project.Title)
	return project, nil
}

func (p *DB) LisProjectByAuthor(ctx context.Context) ([]Project, error) {
	query := "select * from projects order by author_id"

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err = rows.Scan(
			&project.ProjectID,
			&project.AuthorID,
			&project.Title,
			&project.Category,
			&project.Subcategory,
			&project.Description,
			&project.Link,
			&project.Details,
			&project.Payments,
			&project.FundingGoal,
			&project.FundsRaised,
			&project.CreatedAt,
			&project.UpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	p.infol.Printf("projects listed by author: %v", projects)
	return projects, nil
}

func (p *DB) LisProjectByCategory(ctx context.Context) ([]Project, error) {
	query := "select * from projects order by category"

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []Project
	for rows.Next() {
		var project Project
		err = rows.Scan(
			&project.ProjectID,
			&project.AuthorID,
			&project.Title,
			&project.Category,
			&project.Subcategory,
			&project.Description,
			&project.Link,
			&project.Details,
			&project.Payments,
			&project.FundingGoal,
			&project.FundsRaised,
			&project.CreatedAt,
			&project.UpdatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	p.infol.Printf("projects listed by category: %v", projects)
	return projects, nil
}

type NewestPrjRow struct {
	ProjectID   int32
	Title       string
	Category    string
	Subcategory string
	CreatedAt   pgtype.Timestamptz
}

func (p *DB) ListNewestPrj(ctx context.Context) ([]NewestPrjRow, error) {
	query := "SELECT project_id, title, category, sub_category,created_at FROM projects ORDER BY created_at DESC LIMIT 10"

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []NewestPrjRow
	for rows.Next() {
		var prj NewestPrjRow
		err = rows.Scan(
			&prj.ProjectID,
			&prj.Title,
			&prj.Category,
			&prj.Subcategory,
			&prj.CreatedAt)
		if err != nil {
			return nil, err
		}
		projects = append(projects, prj)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	p.infol.Printf("projects listed by category: %v", projects)
	return projects, nil
}
