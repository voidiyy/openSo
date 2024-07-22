package postgres

import (
	"github.com/jackc/pgx/v5/pgtype"
)

//add validator and json

type Author struct {
	AuthorID       int64              `db:"author_id" json:"authorID" validate:"required"`
	NickName       string             `db:"nick_name" json:"nickName" validate:"required"`
	Email          string             `db:"email" json:"email" validate:"required,email"`
	PasswordHash   string             `db:"password_hash" json:"passwordHash" validate:"required"`
	Payments       string             `db:"payments" json:"payments" validate:"required"`
	Bio            string             `db:"bio" json:"bio" validate:"required"`
	Link           string             `db:"link" json:"link" validate:"required"`
	AdditionalInfo string             `db:"additional_info" json:"additionalInfo" validate:"required"`
	CreatedAt      pgtype.Timestamptz `db:"created_at" json:"createdAt" validate:"required"`
	LastLogin      pgtype.Timestamptz `db:"last_login" json:"lastLogin" validate:"required"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at" json:"updatedAt" validate:"required"`
}

type Comment struct {
	Commentator   string             `db:"commentator" json:"commentator" validate:"required"`
	PostType      string             `db:"post_type" json:"postType" validate:"required"`
	CommentatorID int64              `db:"commentator_id" json:"commentatorId" validate:"required"`
	PostID        int64              `db:"post_id" json:"postId" validate:"required"`
	Comment       string             `db:"comment" json:"comment" validate:"required"`
	CommentDate   pgtype.Timestamptz `db:"comment_date" json:"commentDate" validate:"required"`
}

type OrgSupporter struct {
	EntityType     string             `db:"entity_type" json:"entityType" validate:"required"`
	EntityID       int64              `db:"entity_id" json:"entityId" validate:"required"`
	OrgID          int64              `db:"org_id" json:"orgId" validate:"required"`
	DonationAmount pgtype.Numeric     `db:"donation_amount" json:"donationAmount" validate:"required"`
	DonationDate   pgtype.Timestamptz `db:"donation_date" json:"donationDate" validate:"required"`
}

type Organization struct {
	OrgID          int32              `db:"org_id" json:"orgId" validate:"required"`
	AuthorID       int64              `db:"author_id" json:"authorId" validate:"required"`
	Category       string             `db:"category" json:"category" validate:"required"`
	Subcategory    string             `db:"subcategory" json:"subcategory" validate:"required"`
	Name           string             `db:"name" json:"name" validate:"required"`
	Description    string             `db:"description" json:"description" validate:"required"`
	Website        string             `db:"website" json:"website" validate:"required"`
	ContactEmail   string             `db:"contact_email" json:"contactEmail" validate:"required"`
	LogoUrl        string             `db:"logo_url" json:"logoUrl" validate:"required"`
	AdditionalInfo string             `db:"additional_info" json:"additionalInfo" validate:"required"`
	CreatedAt      pgtype.Timestamptz `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt      pgtype.Timestamptz `db:"updated_at" json:"updatedAt" validate:"required"`
}

type Project struct {
	ProjectID   int32              `db:"project_id" json:"projectId" validate:"required"`
	AuthorID    int64              `db:"author_id" json:"authorId" validate:"required"`
	Title       string             `db:"title" json:"title" validate:"required"`
	Category    string             `db:"category" json:"category" validate:"required"`
	Subcategory string             `db:"sub_category" json:"subcategory"`
	Description string             `db:"description" json:"description" validate:"required"`
	Link        string             `db:"link" json:"link" validate:"required"`
	Details     string             `db:"details" json:"details" validate:"required"`
	Payments    string             `db:"payments" json:"payments" validate:"required"`
	Status      bool               `db:"status" json:"status" validate:"required"`
	FundingGoal pgtype.Numeric     `db:"funding_goal" json:"fundingGoal" validate:"required"`
	FundsRaised pgtype.Numeric     `db:"funds_raised" json:"fundsRaised" validate:"required"`
	CreatedAt   pgtype.Timestamptz `db:"created_at" json:"createdAt" validate:"required"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at" json:"updatedAt" validate:"required"`
}

type ProjectSupporter struct {
	EntityType     string             `db:"entity_type" json:"entityType" validate:"required"`
	EntityID       int64              `db:"entity_id" json:"entityId" validate:"required"`
	ProjectID      int64              `db:"project_id" json:"projectId" validate:"required"`
	DonationAmount pgtype.Numeric     `db:"donation_amount" json:"donationAmount" validate:"required"`
	DonationDate   pgtype.Timestamptz `db:"donation_date" json:"donationDate" validate:"required"`
}

type TransactionsCard struct {
	ID              int32              `db:"id" json:"id" validate:"required"`
	UserID          pgtype.Int8        `db:"user_id" json:"userId" validate:"required"`
	AuthorID        pgtype.Int8        `db:"author_id" json:"authorId" validate:"required"`
	SenderAddr      string             `db:"sender_addr" json:"senderAddr" validate:"required"`
	ReceiverAddr    string             `db:"receiver_addr" json:"receiverAddr" validate:"required"`
	ProjectID       pgtype.Int8        `db:"project_id" json:"projectId" validate:"required"`
	Amount          pgtype.Numeric     `db:"amount" json:"amount" validate:"required"`
	TransactionDate pgtype.Timestamptz `db:"transaction_date" json:"transactionDate" validate:"required"`
	PaymentMethod   pgtype.Text        `db:"payment_method" json:"paymentMethod" validate:"required"`
}

type TransactionsCrypto struct {
	ID              int32              `db:"id" json:"id" validate:"required"`
	UserID          pgtype.Int8        `db:"user_id" json:"userId" validate:"required"`
	AuthorID        pgtype.Int8        `db:"author_id" json:"authorId" validate:"required"`
	SenderAddr      string             `db:"sender_addr" json:"senderAddr" validate:"required"`
	ReceiverAddr    string             `db:"receiver_addr" json:"receiverAddr" validate:"required"`
	Network         string             `db:"network" json:"network" validate:"required"`
	Tax             pgtype.Numeric     `db:" tax" json:"tax" validate:"required"`
	ProjectID       pgtype.Int8        `db:"project_id" json:"projectId" validate:"required"`
	Amount          pgtype.Numeric     `db:"amount" json:"amount" validate:"required"`
	TransactionDate pgtype.Timestamptz `db:"transaction_date" json:"transactionDate" validate:"required"`
	PaymentMethod   pgtype.Text        `db:"payment_method" json:"paymentMethod" validate:"required"`
}

type User struct {
	UserID       int64              `db:"user_id" json:"userId" validate:"required"`
	Username     string             `db:"username" json:"username" validate:"required"`
	Email        string             `db:"email" json:"email" validate:"required,email"`
	PasswordHash string             `db:"password_hash" json:"passwordHash" validate:"required"`
	CreatedAt    pgtype.Timestamptz `db:"created_at" json:"createdAt" validate:"required"`
	LastLogin    pgtype.Timestamptz `db:"last_login" json:"lastLogin" validate:"required"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at" json:"updatedAt" validate:"required"`
}
