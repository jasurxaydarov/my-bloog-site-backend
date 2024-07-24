package modles

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ArticleID     uuid.UUID `json:"article_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
	CategoryID    uuid.UUID `json:"category_id"`
	SubCategoryID uuid.UUID `json:"sub_category_id"`
}

// UpdateReq represents the request payload for updating an existing article.
type ArticleReq struct {
	ArticleID     uuid.UUID `json:"article_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CategoryID    string    `json:"category_id"`
	SubCategoryID string    `json:"sub_category_id"`
}

type ArticleResp struct {
	ArticleID     uuid.UUID `json:"article_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CategoryID    uuid.UUID `json:"category_id"`
	SubCategoryID uuid.UUID `json:"sub_category_id"`
}

type UpdateArticleReq struct {
	ArticleID     string    `json:"article_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	CategoryID    string    `json:"category_id"`
	SubCategoryID string    `json:"sub_category_id"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type GetArticleListResp struct {
	Articles []*Article `json:"articles"`
	Count    int32      `json:"count"`
}
