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

type SubCategoriesResp struct {
	Category SubCategory
	Count    int
}

type SubCategoryResp struct {
	Name       string    `json:"name"`
	CategoryID uuid.UUID `json:"category_id"`
}
