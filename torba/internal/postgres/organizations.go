package postgres

import (
	"context"
	"errors"
)

type CreateOrgParams struct {
	AuthorID       int64
	Category       string
	Subcategory    string
	Name           string
	Description    string
	Website        string
	ContactEmail   string
	AdditionalInfo string
}

func (p *DB) CategoryExistsOrg(ctx context.Context, category string) error {
	query := "SELECT EXISTS (SELECT 1 FROM organizations WHERE category = $1)"
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

func (p *DB) CategoryAndSubCategoryExistsOrg(ctx context.Context, category, subcategory string) error {
	query := "SELECT EXISTS (SELECT 1 FROM organizations WHERE category = $1 and sub_category = $2)"

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

func (p *DB) CreateOrganization(ctx context.Context, org *CreateOrgParams) error {
	query := "insert into organizations (author_id, category, sub_category,description, name, description, website, contact_email,additional_info) values ($1, $2,$3,$4,$5,$6,$7,$8,now())"

	_, err := p.pdb.Exec(ctx, query, org.AuthorID, org.Category, org.Subcategory, org.Description, org.Name, org.Description, org.Website, org.ContactEmail, org.AdditionalInfo, org.AdditionalInfo)
	if err != nil {
		p.errl.Printf("could not create organization: %s", err)
		return err
	}

	p.infol.Printf("created organization : %v", org.Name)
	return nil
}

func (p *DB) DeleteOrganization(ctx context.Context, orgID int64) error {
	query := "delete from organizations where org_id = $1"

	_, err := p.pdb.Exec(ctx, query, orgID)
	if err != nil {
		p.errl.Printf("could not delete organization: %s", err)
		return err
	}
	p.infol.Printf("deleted organization : %v", orgID)
	return nil
}

type UpdateOrgParams struct {
	OrgID          int64
	Category       string
	Subcategory    string
	Name           string
	Description    string
	Website        string
	ContactEmail   string
	AdditionalInfo string
}

func (p *DB) UpdateOrg(ctx context.Context, org *UpdateOrgParams) error {
	query := "update organizations set category = $2, sub_category = $3,name = $4, description = $5, website = $6, contact_email = $7, additional_info = $8 where org_id = $1"

	_, err := p.pdb.Exec(ctx, query, org.OrgID, org.Category, org.Subcategory, org.Name, org.Description, org.AdditionalInfo)
	if err != nil {
		p.errl.Printf("could not update organization: %s", err)
		return err
	}

	p.infol.Printf("updated organization : %v", org.Name)
	return nil
}

func (p *DB) GetOrganizationByID(ctx context.Context, orgID int64) (*Organization, error) {
	query := "select * from organizations where org_id = $1"

	row := p.pdb.QueryRow(ctx, query, orgID)
	org := Organization{}
	err := row.Scan(
		&org.OrgID,
		&org.AuthorID,
		&org.Category,
		&org.Subcategory,
		&org.Name,
		&org.Description,
		&org.Website,
		&org.ContactEmail,
		&org.AdditionalInfo,
		&org.CreatedAt,
		&org.UpdatedAt,
	)
	if err != nil {
		p.errl.Printf("could not get organization: %s", err)
		return nil, err
	}
	return &org, nil
}

func (p *DB) GetOrganizationByName(ctx context.Context, name string) (*Organization, error) {
	query := "select * from organizations where name = $1"

	row := p.pdb.QueryRow(ctx, query, name)
	org := Organization{}
	err := row.Scan(
		&org.OrgID,
		&org.AuthorID,
		&org.Category,
		&org.Subcategory,
		&org.Name,
		&org.Description,
		&org.Website,
		&org.ContactEmail,
		&org.AdditionalInfo,
		&org.CreatedAt,
		&org.UpdatedAt,
	)
	if err != nil {
		p.errl.Printf("could not get organization: %s", err)
		return nil, err
	}
	return &org, nil
}

func (p *DB) ListOrganizationsByCategory(ctx context.Context, category string) ([]*Organization, error) {
	query := "select * from organizations where category = $1"

	rows, err := p.pdb.Query(ctx, query, category)
	if err != nil {
		p.errl.Printf("could not list organizations: %s", err)
		return nil, err
	}
	defer rows.Close()
	var orgs []*Organization
	for rows.Next() {
		org := Organization{}
		err := rows.Scan(
			&org.OrgID,
			&org.AuthorID,
			&org.Category,
			&org.Subcategory,
			&org.Name,
			&org.Description,
			&org.Website,
			&org.ContactEmail,
			&org.AdditionalInfo,
			&org.CreatedAt,
			&org.UpdatedAt)
		if err != nil {
			p.errl.Printf("could scan rows of list organizations: %s", err)
			return nil, err
		}
		orgs = append(orgs, &org)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("could not list organizations: %s", err)
		return nil, err
	}
	return orgs, nil
}

func (p *DB) ListOrganizationsByCatSub(ctx context.Context, category, subCategory string) ([]Organization, error) {
	query := "select * from organizations where category = $1 and sub_category = $2"

	rows, err := p.pdb.Query(ctx, query, category, subCategory)
	if err != nil {
		p.errl.Printf("could not list organizations: %s", err)
		return nil, err
	}
	defer rows.Close()

	var orgs []Organization
	for rows.Next() {
		org := Organization{}
		err := rows.Scan(
			&org.OrgID,
			&org.AuthorID,
			&org.Category,
			&org.Subcategory,
			&org.Name,
			&org.Description,
			&org.Website,
			&org.ContactEmail,
			&org.AdditionalInfo,
			&org.CreatedAt,
			&org.UpdatedAt)
		if err != nil {
			p.errl.Printf("could scan rows of list organizations: %s", err)
			return nil, err
		}
		orgs = append(orgs, org)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("could not list organizations: %s", err)
		return nil, err
	}
	return orgs, nil
}

func (p *DB) ListOrganizationsByAuthor(ctx context.Context, authorID int64) ([]Organization, error) {
	query := "select * from organizations where author_id = $1"

	rows, err := p.pdb.Query(ctx, query, authorID)
	if err != nil {
		p.errl.Printf("could not get organization: %s", err)
		return nil, err
	}
	defer rows.Close()

	var orgs []Organization
	for rows.Next() {
		org := Organization{}
		err := rows.Scan(
			&org.OrgID,
			&org.AuthorID,
			&org.Category,
			&org.Subcategory,
			&org.Name,
			&org.Description,
			&org.AdditionalInfo,
			&org.CreatedAt,
			&org.UpdatedAt)
		if err != nil {
			p.errl.Printf("could not get organization: %s", err)
			return nil, err
		}
		orgs = append(orgs, org)
	}
	if err := rows.Err(); err != nil {
		p.errl.Printf("could not get organization: %s", err)
		return nil, err
	}
	return orgs, nil
}
