package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (v *viwerRepo) CreateComment(ctx context.Context, req *modles.CreateCommentReq) (*modles.Comment, error) {
	req.CommentID = uuid.New()

	query := `	
		INSERT INTO
			comments(
				comment_id,
				comment,
				viwer_id,
				article_id
			)VALUES(
				$1,$2,$3,$4
			)`

	_, err := v.db.Exec(
		ctx,
		query,
		req.CommentID,
		req.Comment,
		req.ViewerID,
		req.ArticleID,
	)

	if err != nil {
		v.log.Error("error on Create comment ", logger.Error(err))
		return nil, err

	}
	resp, err := v.GetComment(ctx, req.CommentID.String())

	if err != nil {
		v.log.Error("error on  GetComment", logger.Error(err))
		return nil, err

	}

	return resp, nil
}
func (v *viwerRepo) GetComment(ctx context.Context, id string) (*modles.Comment, error) {
	var resp modles.Comment

	query := `
		SELECT
  		  * 
		FROM 
   			comments 
		WHERE 
  			comment_id = $1;

	`

	err := v.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&resp.CommentID,
		&resp.Comment,
		&resp.CreatedAt,
		&resp.ViewerID,
		&resp.ArticleID,
	)

	if err != nil {
		v.log.Error("error on GetComment", logger.Error(err))
		return nil, err

	}

	v.log.Debug("sucssesfully Get Comment")
	return &resp, nil
}

func (v *viwerRepo) GetComments(ctx context.Context, getList modles.GetList) (*modles.GetCommments, error) {
	var comment modles.Comment
	var resp modles.GetCommments
	var count int32

	offset := (getList.Pge - 1) * getList.Limit

	query := `
		SELECT 
   			*
		FROM
    		comments
		LIMIT $1 OFFSET $2
		`
			

	rows, err := v.db.Query(ctx, query, getList.Limit, offset)

	if err != nil {
		v.log.Error("error on GetComments", logger.Error(err))
		return nil, err

	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&comment.CommentID,
			&comment.Comment,
			&comment.CreatedAt,
			&comment.ViewerID,
			&comment.ArticleID,
		)
		if err != nil {
			v.log.Error("error on GetComments rows", logger.Error(err))
			return nil, err

		}

		resp.Comments = append(resp.Comments, &comment)

	}

	query = `
		SELECT 
			count(*)
		FROM
			comments
	`

	err = v.db.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		v.log.Error("error on GetComments count", logger.Error(err))
		return nil, err

	}

	resp.Count=count

	return &resp, nil
}
func (v *viwerRepo) UpdateComment(ctx context.Context, req modles.CreateCommentReq) (*modles.Comment, error) {

	return nil, nil
}
func (v *viwerRepo) DeleteComment(ctx context.Context, id string) error {

	return nil
}
