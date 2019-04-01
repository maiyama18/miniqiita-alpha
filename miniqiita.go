package miniqiita

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Token      string
	Logger     *log.Logger
}

func New(rawBaseURL, token string, logger *log.Logger) (*Client, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}

	if logger == nil {
		logger = log.New(os.Stderr, "[LOG]", log.LstdFlags)
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Token:      token,
		Logger:     logger,
	}, nil
}

// Item represents an article published on qiita.com
type Item struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	LikesCount int    `json:"likes_count"`
}

// GetUserItems returns items created by the user with provided userID
//
// GET /api/v2/users/:user_id/items
// https://qiita.com/api/v2/docs#get-apiv2usersuser_iditems
func (c *Client) GetUserItems(ctx context.Context, userID string, page, perPage int) ([]*Item, error) {
	return nil, nil
}
