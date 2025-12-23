package article

import (
	"context"

	"github.com/bxcodec/go-clean-arch/domain/GO/domain"
)

type ArticleRepository interface {
	GetByID(ctx context.Context, id int64) (domain.Article, error)
}

type AuthorRepository interface {
	GetByID(ctx context.Context, id int64) (domain.Author, error)
}

type Service struct {
	articleRepo ArticleRepository
	authorRepo  AuthorRepository
}

func NewService(a ArticleRepository, ar AuthorRepository) *Service {
	return &Service{
		articleRepo: a,
		authorRepo:  ar,
	}
}

func (a *Service) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	res, err = a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	resAuthor, err := a.authorRepo.GetByID(ctx, res.Author.ID)
	if err != nil {
		return domain.Article{}, err
	}
	res.Author = resAuthor
	return

}
