package getfeeds

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"math"
	"os"

	"github.com/Josh2604/go-infinite-scroll/api/core/contracts/getposts"
	"github.com/Josh2604/go-infinite-scroll/api/core/entities"
)

type UseCase interface {
	GetPosts(ctx context.Context, paginator *getposts.Paginator) (*getposts.Response, error)
}

type Implementation struct {
}

// GetFeeds -
func (useCase *Implementation) GetPosts(ctx context.Context, paginator *getposts.Paginator) (*getposts.Response, error) {
	var pageNumber, items = paginator.PageNo, paginator.Limit
	posts := getPosts()
	total := len(posts)

	start := (pageNumber - 1) * items
	end := pageNumber * items

	div := float64(total) / float64(items)
	totalPages := math.Trunc(div)

	HASMORE := true

	if (pageNumber + 1) > int(totalPages) {
		HASMORE = false
	}

	if (paginator.PageNo * paginator.Limit) > total {
		start = 0
		end = 0
	}

	response := getposts.Response{
		Total:       total,
		CurrentPage: pageNumber,
		PagesNo:     int(totalPages),
		HasMore:     HASMORE,
		Items:       posts[start:end],
	}
	return &response, nil
}

func getPosts() []entities.Post {
	posts := make([]entities.Post, 100)

	raw, err := ioutil.ReadFile("feeds.json")
	if err != nil {
		os.Exit(1)
	}
	errJ := json.Unmarshal(raw, &posts)
	if errJ != nil {
		os.Exit(1)
	}
	return posts
}
