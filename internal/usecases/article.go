package usecases

import (
	"github.com/y3kawaguchi/knowledge/internal/domains"
	"github.com/y3kawaguchi/knowledge/internal/repositories"
)

// ArticleUsecase ...
type ArticleUsecase struct {
	repository repositories.ArticleRepository
}

// NewArticleUsecase ...
func NewArticleUsecase(repository *repositories.ArticleRepository) *ArticleUsecase {
	usecase := &ArticleUsecase{
		repository: *repository,
	}
	return usecase
}

// Create ...
func (a *ArticleUsecase) Create(article *domains.Article) (int64, error) {
	return a.repository.Save(article)
}

// Get ...
func (a *ArticleUsecase) Get() (*domains.Articles, error) {
	articles, err := a.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return articles, nil
}
