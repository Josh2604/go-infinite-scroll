package dependencies

import (
	"github.com/Josh2604/go-infinite-scroll/api/entrypoints"
	"github.com/Josh2604/go-infinite-scroll/api/usecases/getfeeds"
)

type Handlers struct {
	GetPosts entrypoints.Handler
}

func Exec() *Handlers {
	// UseCases
	postsUseCases := &getfeeds.Implementation{}

	// Handlers
	handlers := Handlers{}

	handlers.GetPosts = &entrypoints.GetPosts{
		GetPostsUseCase: postsUseCases,
	}

	return &handlers
}
