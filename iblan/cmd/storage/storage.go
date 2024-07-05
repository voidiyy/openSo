package storage

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"iblan/cmd/structures"
	"log"
	"os"
)

/*func (s *PostgresStore) DBInit() ([]error, error) {
	var errs []error

	err := s.createUserTable()
	if err != nil {
		errs := append(errs, err)
		return errs, nil
	}

	err = s.createMemberTable()
	if err != nil {
		errs := append(errs, err)
		return errs, nil
	}

	err = s.createArticleTable()
	if err != nil {
		errs := append(errs, err)
		return errs, nil
	}
	return errs, nil
}*/

var Instance *gorm.DB

type GlobalStorage interface {
	UserStorage
	MemberStorage
	StorageForArticles
}

type PostgresStore struct {
	db *gorm.DB
}

func NewStorage() (*PostgresStore, error) {
	godotenv.Load(".env")

	connStr := os.Getenv("DB_STR")
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	Instance = db
	return &PostgresStore{db: db}, nil
}

func (*PostgresStore) Migrate() error {
	if Instance == nil {
		return errors.New("Database instance is not initialized")
	}
	err := Instance.AutoMigrate(&structures.User{}, &structures.Article{}, &structures.Member{})
	if err != nil {
		return fmt.Errorf("ddss", err)
	}
	return nil
}
