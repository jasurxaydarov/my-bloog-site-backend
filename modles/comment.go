package modles

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID `json:"comment_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	ViewerID  uuid.UUID `json:"viewer_id"`
	ArticleID uuid.UUID `json:"article_id"`
}

type CreateCommentReq struct {
	CommentID uuid.UUID `json:"comment_id"`
	Comment   string    `json:"comment"`
	ViewerID  uuid.UUID `json:"viewer_id"`
	ArticleID uuid.UUID `json:"article_id"`
}
type UpdateCommentReq struct {
	CommentID uuid.UUID `json:"comment_id"`
	Comment   string    `json:"comment"`
	ArticleID uuid.UUID `json:"article_id"`
}
type GetCommments struct {
	Comments []*Comment `json:"comments"`
	Count    int32      `json:"count"`
}

