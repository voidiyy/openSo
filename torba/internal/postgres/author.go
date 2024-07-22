package postgres

import (
	"context"
	"errors"
)

func (p *DB) AuthorExists(ctx context.Context, nickname string) error {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM authors WHERE nick_name=$1)"

	err := p.pdb.QueryRow(ctx, query, nickname).Scan(&exists)
	if err != nil {
		return errors.New("error checking for existence of author")
	}

	if exists {
		return errors.New("author already exists")
	}
	return nil
}

func (p *DB) EmailExistsA(ctx context.Context, email string) error {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM authors WHERE email=$1)"

	err := p.pdb.QueryRow(ctx, query, email).Scan(&exists)
	if err != nil {
		return errors.New("error checking for existence of email")
	}
	if exists {
		return errors.New("email already exists")
	}
	return nil
}

type CreateAuthorParams struct {
	NickName     string `validate:"required"`
	Email        string `validate:"required"`
	PasswordHash string `validate:"required"`
}

func (p *DB) CreateAuthor(ctx context.Context, arg *CreateAuthorParams) error {
	query := "INSERT INTO authors (nick_name, email, password_hash, created_at) VALUES ($1, $2, $3, now())"

	_, err := p.pdb.Exec(ctx, query, arg.NickName, arg.Email, arg.PasswordHash)
	if err != nil {
		p.errl.Printf("failed to insert author: %v", err)
		return err
	}

	p.infol.Printf("author created: %v", arg.NickName)
	return nil
}

type CreateFullAuthorParams struct {
	NickName       string `validate:"required"`
	Email          string `validate:"required"`
	PasswordHash   string `validate:"required"`
	Payments       string `validate:"required"`
	Bio            string `validate:"required"`
	Link           string `validate:"required"`
	AdditionalInfo string `validate:"required"`
}

func (p *DB) CreateFullAuthor(ctx context.Context, arg *CreateFullAuthorParams) error {
	query := "insert into authors(nick_name, email, password_hash, payments, bio, link, additional_info, updated_at) values ($1, $2, $3,$4,$5,$6,$7,now())"

	_, err := p.pdb.Exec(ctx, query, arg.NickName, arg.Email, arg.PasswordHash, arg.Payments, arg.Bio, arg.Link, arg.AdditionalInfo)
	if err != nil {
		p.errl.Printf("failed to insert full author: %v", err)
		return err
	}

	p.infol.Printf("author created: %v", arg.NickName)
	return nil
}

func (p *DB) DeleteAuthor(ctx context.Context, id int64) error {
	query := "DELETE FROM authors WHERE author_id=$1"
	_, err := p.pdb.Exec(ctx, query, id)
	if err != nil {
		p.errl.Printf("failed to delete author: %v", err)
		return err
	}

	p.infol.Printf("author deleted: %v", id)
	return nil
}

func (p *DB) GetAuthorByID(ctx context.Context, id int64) (Author, error) {
	query := "SELECT * FROM authors WHERE author_id = $1 LIMIT 1"

	row := p.pdb.QueryRow(ctx, query, id)
	var a Author
	err := row.Scan(
		&a.AuthorID,
		&a.NickName,
		&a.Email,
		&a.PasswordHash,
		&a.Payments,
		&a.Bio,
		&a.Link,
		&a.AdditionalInfo,
		&a.CreatedAt,
		&a.LastLogin,
		&a.UpdatedAt,
	)

	p.infol.Printf("author goted by ID: %v", a)
	return a, err
}

func (p *DB) GetAuthorByName(ctx context.Context, nick string) (Author, error) {
	query := "SELECT * FROM authors WHERE nick_name = $1 LIMIT 1"

	row := p.pdb.QueryRow(ctx, query, nick)
	var a Author
	err := row.Scan(
		&a.AuthorID,
		&a.NickName,
		&a.Email,
		&a.PasswordHash,
		&a.Payments,
		&a.Bio,
		&a.Link,
		&a.AdditionalInfo,
		&a.CreatedAt,
		&a.LastLogin,
		&a.UpdatedAt,
	)
	p.infol.Printf("author goted by name: %v", a)
	return a, err
}

func (p *DB) ListAuthorsByID(ctx context.Context) ([]Author, error) {
	query := "SELECT * FROM authors order by author_id"

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		p.errl.Printf("failed to list authors by ID: %v", err)
		return nil, err
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		a := Author{}
		err = rows.Scan(
			&a.AuthorID,
			&a.NickName,
			&a.Email,
			&a.PasswordHash,
			&a.Payments,
			&a.Bio,
			&a.Link,
			&a.AdditionalInfo,
			&a.CreatedAt,
			&a.LastLogin,
			&a.UpdatedAt)
		if err != nil {
			p.errl.Printf("failed to scan authors by ID: %v", err)
			return nil, err
		}
		authors = append(authors, a)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("failed to list authors by ID: %v", err)
		return nil, err
	}

	p.infol.Printf("listed authors by ID: %v", authors)
	return authors, nil
}

func (p *DB) ListAuthorsByNick(ctx context.Context) ([]Author, error) {
	query := "SELECT * FROM authors order by nick_name"

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		p.errl.Printf("failed to list authors by Nick: %v", err)
		return nil, err
	}
	defer rows.Close()

	var authors []Author
	for rows.Next() {
		var a Author
		err = rows.Scan(
			&a.AuthorID,
			&a.NickName,
			&a.Email,
			&a.PasswordHash,
			&a.Payments,
			&a.Bio,
			&a.Link,
			&a.AdditionalInfo,
			&a.CreatedAt,
			&a.LastLogin,
			&a.UpdatedAt)
		if err != nil {
			p.errl.Printf("failed to scan authors by Nick: %v", err)
			return nil, err
		}
		authors = append(authors, a)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("failed to list authors by Nick: %v", err)
		return nil, err
	}

	p.infol.Printf("listed authors by Nick: %v", authors)
	return authors, nil
}

type UpdateAuthorFullParams struct {
	AuthorID       int64  `validate:"required"`
	NickName       string `validate:"required"`
	Email          string `validate:"required"`
	PasswordHash   string `validate:"required"`
	Payments       string `validate:"required"`
	Bio            string `validate:"required"`
	Link           string `validate:"required"`
	AdditionalInfo string `validate:"required"`
}

func (p *DB) UpdateAuthor(ctx context.Context, arg *UpdateAuthorFullParams) error {
	query := "UPDATE authors set nick_name = $2, email = $3, password_hash = $4, payments = $5, bio = $6, link = $7, additional_info = $9 WHERE author_id = $1"

	_, err := p.pdb.Exec(ctx, query, arg.AuthorID,
		arg.NickName,
		arg.Email,
		arg.PasswordHash,
		arg.Payments,
		arg.Bio,
		arg.Link,
		arg.AdditionalInfo)
	if err != nil {
		p.errl.Printf("failed to update author: %v", err)
		return err
	}

	p.infol.Printf("author updated: %v", arg.AuthorID)
	return nil
}

//advanced queries

type ListOrgRow struct {
	OrgID          int32
	Name           string
	Category       string
	SupporterCount int64
	TotalDonations int64
}

func (p *DB) ListOrgByAuthor(ctx context.Context, id int64) ([]ListOrgRow, error) {
	query := "SELECT o.org_id, o.name, o.category, COUNT(os.entity_id) AS supporter_count, SUM(os.donation_amount) AS total_donations FROM organizations o LEFT JOIN org_supporters os ON o.org_id = os.org_id WHERE o.author_id = $1 GROUP BY o.org_id"

	rows, err := p.pdb.Query(ctx, query, id)
	if err != nil {
		p.errl.Printf("failed to list org by author: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orgs []ListOrgRow
	for rows.Next() {
		var o ListOrgRow
		err = rows.Scan(
			&o.OrgID,
			&o.Name,
			&o.Category,
			&o.SupporterCount,
			&o.TotalDonations)
		if err != nil {
			p.errl.Printf("failed to list org by author: %v", err)
			return nil, err
		}
		orgs = append(orgs, o)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("failed to list org by author: %v", err)
		return nil, err
	}

	p.infol.Printf("listed orgs by author: %v", orgs)
	return orgs, nil
}

type ListProjectsRow struct {
	ProjectID      int32
	Title          string
	Category       string
	SupporterCount int64
	TotalDonations int64
}

func (p *DB) ListProjectsByAuthor(ctx context.Context, id int64) ([]ListProjectsRow, error) {
	query := "SELECT p.project_id, p.title, p.category, COUNT(ps.entity_id) AS supporter_count, SUM(ps.donation_amount) AS total_donations FROM projects p LEFT JOIN project_supporters ps ON p.project_id = ps.project_id WHERE p.author_id = $1 GROUP BY p.project_id"

	rows, err := p.pdb.Query(ctx, query, id)
	if err != nil {
		p.errl.Printf("failed to list projects by author: %v", err)
		return nil, err
	}
	defer rows.Close()

	var projects []ListProjectsRow
	for rows.Next() {
		var o ListProjectsRow
		err = rows.Scan(
			&o.ProjectID,
			&o.Title,
			&o.Category,
			&o.SupporterCount,
			&o.TotalDonations)
		if err != nil {
			p.errl.Printf("failed to list projects by author: %v", err)
			return nil, err
		}
		projects = append(projects, o)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("failed to list projects by author: %v", err)
		return nil, err
	}

	p.infol.Printf("listed projects by author: %v", projects)
	return projects, nil
}

type ListAllRow struct {
	AuthorID           int64
	NickName           string
	TotalProjects      int64
	TotalOrganizations int64
}

func (p *DB) ListAll(ctx context.Context) ([]ListAllRow, error) {
	query := "SELECT a.author_id, a.nick_name, COUNT(DISTINCT p.project_id) AS total_projects, COUNT(DISTINCT o.org_id) AS total_organizations FROM authors a LEFT JOIN projects p ON a.author_id = p.author_id LEFT JOIN organizations o ON a.author_id = o.author_id GROUP BY a.author_id"

	rows, err := p.pdb.Query(ctx, query)
	if err != nil {
		p.errl.Printf("failed to list authors: %v", err)
		return nil, err
	}

	defer rows.Close()

	var list []ListAllRow

	for rows.Next() {
		var l ListAllRow
		err := rows.Scan(
			&l.AuthorID,
			&l.NickName,
			&l.TotalProjects,
			&l.TotalOrganizations,
		)
		if err != nil {
			p.errl.Printf("failed to list authors: %v", err)
			return nil, err
		}
		list = append(list, l)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("failed to list authors: %v", err)
		return nil, err
	}

	p.infol.Printf("listed authors: %v", list)
	return list, nil
}
