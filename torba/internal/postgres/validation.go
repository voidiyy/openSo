package postgres

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"strings"
)

func (p *DB) HashPass(str string) string {

	passwd := []byte(str)

	pass, _ := bcrypt.GenerateFromPassword(passwd, 8)

	return string(pass)
}

func (p *DB) CreateUserValidator(ctx context.Context, usr *CreateUserParams) error {

	username := strings.TrimSpace(usr.Username)
	ru := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	if !ru.MatchString(username) {
		return errors.New("username must only contain letters, numbers, and hyphens")
	}
	err := p.UserExists(ctx, usr.Username)
	if err != nil {
		return errors.New("username already exists")
	}

	email := strings.TrimSpace(usr.Email)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must only contain letters, numbers, and hyphens")
	}
	err = p.EmailExistsU(ctx, email)
	if err != nil {
		return errors.New("email already exists")
	}

	password := strings.TrimSpace(usr.PasswordHash)
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func (p *DB) UpdateUserValidator(ctx context.Context, usr *UpdateUserParams) error {

	username := strings.TrimSpace(usr.Username)
	ru := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	if !ru.MatchString(username) {
		return errors.New("username must only contain letters, numbers, and hyphens")
	}
	err := p.UserExists(ctx, usr.Username)
	if err != nil {
		return errors.New("username already exists")
	}

	email := strings.TrimSpace(usr.Email)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must only contain letters, numbers, and hyphens")
	}

	password := strings.TrimSpace(usr.PasswordHash)
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func (p *DB) CreateFullAuthorValidator(ctx context.Context, author *CreateFullAuthorParams) error {
	authorname := strings.TrimSpace(author.NickName)
	ru := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	if !ru.MatchString(authorname) {
		return errors.New("username must only contain letters, numbers, and hyphens")
	}
	err := p.AuthorExists(ctx, authorname)
	if err != nil {
		return errors.New("author does not exist")
	}

	email := strings.TrimSpace(author.Email)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must only contain letters, numbers, and hyphens")
	}
	err = p.EmailExistsU(ctx, email)
	if err != nil {
		return errors.New("email already exists")
	}

	password := strings.TrimSpace(author.PasswordHash)
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	bio := strings.TrimSpace(author.Bio)
	if len(bio) < 20 {
		return errors.New("bio must at least 20 characters")
	}

	link := strings.TrimSpace(author.Link)
	rl := regexp.MustCompile(`^https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\S*)?$`)

	if !rl.MatchString(link) {
		return errors.New("invalid link")
	}

	addInfo := strings.TrimSpace(author.AdditionalInfo)
	if len(addInfo) < 20 {
		return errors.New("additional info must be at least 20 characters")
	}

	return nil
}

func (p *DB) UpdateAuthorValidator(ctx context.Context, author *UpdateAuthorFullParams) error {
	authorname := strings.TrimSpace(author.NickName)
	ru := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	if !ru.MatchString(authorname) {
		return errors.New("username must only contain letters, numbers, and hyphens")
	}
	err := p.AuthorExists(ctx, authorname)
	if err != nil {
		return errors.New("author does not exist")
	}

	email := strings.TrimSpace(author.Email)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must only contain letters, numbers, and hyphens")
	}
	err = p.EmailExistsU(ctx, email)
	if err != nil {
		return errors.New("email already exists")
	}

	password := strings.TrimSpace(author.PasswordHash)
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	bio := strings.TrimSpace(author.Bio)
	if len(bio) < 20 {
		return errors.New("bio must at least 20 characters")
	}

	link := strings.TrimSpace(author.Link)
	rl := regexp.MustCompile(`^https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\S*)?$`)

	if !rl.MatchString(link) {
		return errors.New("invalid link")
	}

	addInfo := strings.TrimSpace(author.AdditionalInfo)
	if len(addInfo) < 20 {
		return errors.New("additional info must be at least 20 characters")
	}

	return nil
}

func (p *DB) CreateProjectValidator(ctx context.Context, project *CreateProjectParams) error {

	if project.AuthorID <= 0 {
		return errors.New("invalid author id")
	}

	_, err := p.GetAuthorByID(ctx, project.AuthorID)
	if err != nil {
		return errors.New("author not found")
	}

	if len(project.Title) <= 5 {
		return errors.New("project title to short")
	}

	category := strings.TrimSpace(project.Category)

	if len(category) <= 5 {
		return errors.New("category must at least 5 characters")
	}
	err = p.CategoryExistsPrj(ctx, category)
	if err != nil {
		return errors.New("category does not exist")
	}

	subcategory := strings.TrimSpace(project.Subcategory)

	if len(subcategory) <= 5 {
		return errors.New("subcategory must at least 5 characters")
	}
	err = p.CategoryAndSubCategoryExistsPrj(ctx, category, subcategory)
	if err != nil {
		return errors.New("category and subcategory does not exist")
	}

	if len(project.Description) <= 20 {
		return errors.New("description must at least 20 characters")
	}

	link := strings.TrimSpace(project.Link)
	rl := regexp.MustCompile(`^https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\S*)?$`)

	if !rl.MatchString(link) {
		return errors.New("invalid link")
	}

	if len(project.Details) < 20 {
		return errors.New("details must at least 20 characters")
	}

	if len(project.Payments) < 20 {
		return errors.New(" invalid payments")
	}

	re := regexp.MustCompile(`^\d{1,8}(\.\d{1,2})?$`)
	if !re.MatchString(project.FundingGoal) {
		return errors.New("invalid funding goal")
	}

	return nil
}

func (p *DB) UpdateProjectValidator(ctx context.Context, project *UpdateProjectParams) error {

	if project.ProjectID <= 0 {
		return errors.New("invalid project id")
	}

	_, err := p.GetProjectByID(ctx, project.ProjectID)
	if err != nil {
		return errors.New("project with this id not exists")
	}

	if len(project.Title) <= 5 {
		return errors.New("project title to short")
	}

	category := strings.TrimSpace(project.Category)
	if len(category) <= 5 {
		return errors.New("category must at least 5 characters")
	}
	err = p.CategoryExistsPrj(ctx, category)
	if err != nil {
		return errors.New("category does not exist")
	}

	subcategory := strings.TrimSpace(project.Subcategory)
	if len(subcategory) <= 5 {
		return errors.New("subcategory must at least 5 characters")
	}
	err = p.CategoryAndSubCategoryExistsPrj(ctx, category, subcategory)
	if err != nil {
		return errors.New("category and subcategory does not exist")
	}

	if len(project.Description) <= 20 {
		return errors.New("description must at least 20 characters")
	}

	link := strings.TrimSpace(project.Link)
	rl := regexp.MustCompile(`^https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\S*)?$`)

	if !rl.MatchString(link) {
		return errors.New("invalid link")
	}

	if len(project.Details) < 20 {
		return errors.New("details must at least 20 characters")
	}
	if len(project.Payments) < 20 {
		return errors.New(" invalid payments")
	}

	re := regexp.MustCompile(`^\d{1,8}(\.\d{1,2})?$`)
	if !re.MatchString(project.FundingGoal) {
		return errors.New("invalid funding goal")
	}

	return nil
}

func (p *DB) CreateOrgValidator(ctx context.Context, org *CreateOrgParams) error {

	if org.AuthorID <= 0 {
		return errors.New("invalid author id")
	}

	_, err := p.GetAuthorByID(ctx, org.AuthorID)
	if err != nil {
		return errors.New("author not found")
	}

	category := strings.TrimSpace(org.Category)
	if len(category) <= 5 {
		return errors.New("category must at least 5 characters")
	}
	err = p.CategoryExistsOrg(ctx, category)
	if err != nil {
		return errors.New("category does not exist")
	}

	subcategory := strings.TrimSpace(org.Subcategory)
	if len(subcategory) <= 5 {
		return errors.New("subcategory must at least 5 characters")
	}
	err = p.CategoryAndSubCategoryExistsOrg(ctx, category, subcategory)
	if err != nil {
		return errors.New("category and subcategory does not exist")
	}

	name := strings.TrimSpace(org.Name)
	if len(name) <= 5 {
		return errors.New("name must at least 5 characters")
	}

	if len(org.Description) <= 20 {
		return errors.New("description must at least 20 characters")
	}

	link := strings.TrimSpace(org.Website)
	rl := regexp.MustCompile(`^https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\S*)?$`)

	if !rl.MatchString(link) {
		return errors.New("invalid link")
	}

	email := strings.TrimSpace(org.ContactEmail)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must only contain letters, numbers, and hyphens")
	}

	info := strings.TrimSpace(org.AdditionalInfo)
	if len(info) <= 20 {
		return errors.New("additionalInfo must be at lest 20 characters")
	}

	return nil
}

func (p *DB) UpdateOrgValidator(ctx context.Context, org *UpdateOrgParams) error {

	if org.OrgID <= 0 {
		return errors.New("invalid project id")
	}

	_, err := p.GetOrganizationByID(ctx, org.OrgID)
	if err != nil {
		return errors.New("project with this id not exists")
	}

	category := strings.TrimSpace(org.Category)
	if len(category) <= 5 {
		return errors.New("category must at least 5 characters")
	}
	err = p.CategoryExistsOrg(ctx, category)
	if err != nil {
		return errors.New("category does not exist")
	}

	subcategory := strings.TrimSpace(org.Subcategory)
	if len(subcategory) <= 5 {
		return errors.New("subcategory must at least 5 characters")
	}
	err = p.CategoryAndSubCategoryExistsOrg(ctx, category, subcategory)
	if err != nil {
		return errors.New("category and subcategory does not exist")
	}

	name := strings.TrimSpace(org.Name)
	if len(name) < 5 {
		return errors.New("organization name must be at lest 5 characters")
	}

	if len(org.Description) <= 20 {
		return errors.New("description must at least 20 characters")
	}

	link := strings.TrimSpace(org.Website)
	rl := regexp.MustCompile(`^https://[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}(\S*)?$`)

	if !rl.MatchString(link) {
		return errors.New("invalid link")
	}

	email := strings.TrimSpace(org.ContactEmail)
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("email must only contain letters, numbers, and hyphens")
	}

	info := strings.TrimSpace(org.AdditionalInfo)
	if len(info) <= 20 {
		return errors.New("additionalInfo must be at lest 20 characters")
	}
	return nil
}
