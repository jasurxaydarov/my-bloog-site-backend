package modles

import (
	"time"

	"github.com/google/uuid"
)

// /
type Category struct {
	CategoryID uuid.UUID `json:"category_id"`
	Name       string    `json:"name"`
	CreatedAt  time.Time `json:"created_at"`
}

type CategoryResp struct {
	Name string `json:"name"`
}

type GetCategories struct {
	Categories []*Category
	Count      int32
}
