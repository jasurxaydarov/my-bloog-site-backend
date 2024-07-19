package modles

import (
	"time"

	"github.com/google/uuid"
)

type Owner struct {
	FullName    string `json:"full_name"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	LinkedIn    string `json:"linked_in"`
	Telegram    string `json:"telegram"`
	Github      string `json:"github"`
	Leetcode    string `json:"leetcode"`
	AboutMe     string `json:"about_me"`
}

type LoginOwn struct {
	PhoneNumber string `json:"phone_number"`
	Gmail       string `json:"gmail"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

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

// //
type Article struct {
	ArticleID     uuid.UUID  `json:"article_id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
	CategoryID    uuid.UUID  `json:"category_id"`
	SubCategoryID uuid.UUID  `json:"sub_category_id"`
}

// /
type Viewer struct {
	ViewerID uuid.UUID `json:"viewer_id"`
	FullName string    `json:"full_name"`
	Username string    `json:"username"`
	Gmail    string    `json:"gmail"`
	Password string    `json:"password"`
}

type CheackViwer struct {
	Gmail string `json:"gmail"`
}

// ///
type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	ViewerID  uuid.UUID `json:"viewer_id"`
	ArticleID uuid.UUID `json:"article_id"`
}

type GetList struct {
	Limit int32
	Pge   int32
}

type Common struct {
	TableName  string `json:"tanle_name"`
	ColumnName string `json:"clomn_name"`
	ExpValue   any    `json:"exp_value"`
}

type CheckExists struct {
	IsExists bool
	Status   string
}
