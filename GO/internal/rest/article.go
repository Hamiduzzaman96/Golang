package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/bxcodec/go-clean-arch/domain/GO/domain"
)

type ResponseError struct {
	Message string `json:"message"`
}

type ArticleService interface {
	GetByID(ctx context.Context, id int64) (domain.Article, error)
}

// ArticleHandler represents the HTTP handler for articles
type ArticleHandler struct {
	Service ArticleService
}

func NewArticleHandler(e *echo.Echo, svc ArticleService) {
	handler := &ArticleHandler{
		Service: svc,
	}
	e.GET("/articles/:id", handler.GetByID)
}

func (a *ArticleHandler) GetByID(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()

	art, err := a.Service.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)
}
