package modles

import (
	"time"

	"github.com/google/uuid"
)

// /
type SubCategory struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CreatedAt     time.Time `json:"created_at"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type SubCategoryReq struct {
	SubCategoryID uuid.UUID `json:"sub_category_id"`
	Name          string    `json:"name"`
	CategoryID    uuid.UUID `json:"category_id"`
}

type GetSubCategoriesLidtResp struct {
	SubCategory []*SubCategory `json:"sub_categories"`
	Count    int32           `json:"count"`
}
