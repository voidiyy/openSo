package storage

import (
	"fmt"
	"iblan/cmd/structures"
)

type StorageForArticles interface {
	CreateArticle(*structures.Article) error
	UpdateArticle(id int, title, category, body, payments, link string) error
	DeleteArticle(id int) error
	GetArticles() ([]*structures.Article, error)
	GetArticleByID(int) (*structures.Article, error)
	GetArticlesByCategory(string) ([]*structures.Article, error)
	GetArticleFull(string, int) (*structures.Article, error)
}

func (s *PostgresStore) CreateArticle(a *structures.Article) error {
	if err := s.db.Create(a).Error; err != nil {
		return fmt.Errorf("create article error: %w", err)
	}
	fmt.Printf("Article '%v' created successfully\n", a.Title)
	return nil
}

func (s *PostgresStore) UpdateArticle(id int, title, category, body, payments, link string) error {
	article := &structures.Article{}
	if err := s.db.First(article, id).Error; err != nil {
		return fmt.Errorf("update article error: %w", err)
	}
	article.Title = title
	article.Category = category
	article.Body = body
	article.Payments = payments
	article.Link = link
	if err := s.db.Save(article).Error; err != nil {
		return fmt.Errorf("update article error: %w", err)
	}
	fmt.Printf("Article '%v' updated successfully\n", article.Title)
	return nil
}

func (s *PostgresStore) DeleteArticle(id int) error {
	if err := s.db.Delete(&structures.Article{}, id).Error; err != nil {
		return fmt.Errorf("delete article error: %w", err)
	}
	fmt.Printf("Article '%v' deleted successfully\n", id)
	return nil
}

func (s *PostgresStore) GetArticles() ([]*structures.Article, error) {
	articles := []*structures.Article{}
	if err := s.db.Find(&articles).Error; err != nil {
		return nil, fmt.Errorf("error getting articles: %v", err)
	}
	return articles, nil
}

func (s *PostgresStore) GetArticleByID(id int) (*structures.Article, error) {
	article := &structures.Article{}
	if err := s.db.First(article, id).Error; err != nil {
		return nil, fmt.Errorf("error getting article: %v", err)
	}
	return article, nil
}

func (s *PostgresStore) GetArticlesByCategory(category string) ([]*structures.Article, error) {
	articles := []*structures.Article{}
	if err := s.db.Where("category = ?", category).Find(&articles).Error; err != nil {
		return nil, fmt.Errorf("error getting articles: %v", err)
	}
	return articles, nil
}

func (s *PostgresStore) GetArticleFull(category string, id int) (*structures.Article, error) {
	article := structures.Article{}
	if err := s.db.First(&article).Where("category = ? AND id = ?", category, id).Error; err != nil {
		return nil, fmt.Errorf("error getting article: %v", err)
	}
	return &article, nil
}
