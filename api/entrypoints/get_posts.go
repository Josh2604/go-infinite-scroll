package entrypoints

import (
	"net/http"

	"github.com/Josh2604/go-infinite-scroll/api/core/contracts/getposts"
	apiErrors "github.com/Josh2604/go-infinite-scroll/api/errors"
	"github.com/Josh2604/go-infinite-scroll/api/usecases/getfeeds"
	"github.com/gin-gonic/gin"
)

type GetPosts struct {
	GetPostsUseCase getfeeds.UseCase
}

func (useCase *GetPosts) Handle(c *gin.Context) {
	err := useCase.handle(c)
	if err != nil {
		c.JSON(err.Status, err)
	}
}

func (useCase *GetPosts) handle(c *gin.Context) *apiErrors.Error {
	var request getposts.Paginator
	errq := c.BindJSON(&request)
	if errq != nil {
		return apiErrors.NewBadRequest("Invalid Request Parameters", errq.Error())
	}

	response, err := useCase.GetPostsUseCase.GetPosts(c, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Error!")
		return nil
	}

	c.JSON(http.StatusOK, &response)
	return nil
}
