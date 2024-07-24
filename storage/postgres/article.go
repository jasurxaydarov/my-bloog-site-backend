package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (c *ContentRepo) CreateArticle(ctx context.Context, req *modles.ArticleReq) (*modles.ArticleResp, error) {

	req.ArticleID = uuid.New()

	query := `	
		INSERT INTO
			articles(
				article_id,
				title,
				content,
				category_id,
				sub_category_id
			)VALUES(
				$1,$2,$3,$4,$5
			)`

	_, err := c.db.Exec(
		ctx,
		query,
		req.ArticleID,
		req.Title,
		req.Content,
		req.CategoryID,
		req.SubCategoryID,
	)

	if err != nil {
		c.log.Error("error on create article ", logger.Error(err))
		return nil, err

	}
	resp, err := c.GetArticle(ctx, req.ArticleID.String())

	if err != nil {
		c.log.Error("error on  GetCategory", logger.Error(err))
		return nil, err

	}

	return resp, nil
}

func (c *ContentRepo) GetArticle(ctx context.Context, id string) (*modles.ArticleResp, error) {

	var resp modles.ArticleResp
	query := `
		SELECT
  			article_id,
			title,
			content,
			created_at,
			category_id,
			sub_category_id
		FROM 
   			articles 
		WHERE 
  			article_id = $1;

	`

	err := c.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&resp.ArticleID,
		&resp.Title,
		&resp.Content,
		&resp.CreatedAt,
		&resp.CategoryID,
		&resp.SubCategoryID,
	)

	if err != nil {
		c.log.Error("error on article", logger.Error(err))
		return nil, err

	}

	c.log.Debug("sucssesfully Get article")
	return &resp, nil
}

func (c *ContentRepo) GetArticles(ctx context.Context, limit, page int32) (*modles.GetArticleListResp, error) {
	var article modles.Article
	var resp modles.GetArticleListResp
	var count int32

	offset := (page - 1) * limit

	query := `
		SELECT 
			article_id,
			title,
			content,
			created_at,
			category_id,
			sub_category_id
		FROM
			articles
			WHERE 
			deleted_at IS NULL
	  	LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(ctx, query, limit, offset)

	if err != nil {
		c.log.Error("error on articles", logger.Error(err))
		return nil, err

	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&article.ArticleID,
			&article.Title,
			&article.Content,
			&article.CreatedAt,
			&article.CategoryID,
			&article.SubCategoryID,
		)
		if err != nil {
			c.log.Error("error on articles rows", logger.Error(err))
			return nil, err

		}

		resp.Articles = append(resp.Articles, &article)

	}

	query = `
		SELECT 
			count(*)
		FROM
			 articles
	`

	err = c.db.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		c.log.Error("error on articles count", logger.Error(err))
		return nil, err

	}

	resp.Count = count

	return &resp, nil
}
func (c *ContentRepo) UpdateArticle(ctx context.Context, req modles.ArticleReq) error {

	return nil
}
func (c *ContentRepo) DeleteArticle(ctx context.Context, id string) error {

	return nil
}
