package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (c *ContentRepo) CreateCategory(ctx context.Context, req *modles.Category) (*modles.Category, error) {

	req.CategoryID = uuid.New()

	query := `	
		INSERT INTO
			categories(
				category_id,
				name
			)VALUES(
				$1,$2
			)`

	_, err := c.db.Exec(
		ctx,
		query,
		req.CategoryID,
		req.Name,
	)

	if err != nil {
		c.log.Error("error on Createategory ", logger.Error(err))
		return nil, err

	}
	resp, err := c.GetCategory(ctx, req.CategoryID.String())

	if err != nil {
		c.log.Error("error on  GetCategory", logger.Error(err))
		return nil, err

	}

	return resp, nil
}
func (c *ContentRepo) GetCategory(ctx context.Context, id string) (*modles.Category, error) {

	var resp modles.Category
	query := `
		SELECT
  		  * 
		FROM 
   			categories 
		WHERE 
  			category_id = $1;

	`

	err := c.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&resp.CategoryID,
		&resp.Name,
		&resp.CreatedAt,
	)

	if err != nil {
		c.log.Error("error on GetCategory", logger.Error(err))
		return nil, err

	}

	c.log.Debug("sucssesfully Get category")
	return &resp, nil
}
func (c *ContentRepo) GetCategories(ctx context.Context, getList modles.GetList) (*modles.GetCategories, error) {
	var category modles.Category
	var resp modles.GetCategories
	var count int32

	offset := (getList.Pge - 1) * getList.Limit

	query := `
		SELECT 
   			*
		FROM
    		categories
		WHERE 
  			deleted_at IS NULL
		LIMIT $1 OFFSET $2
		`
			

	rows, err := c.db.Query(ctx, query, getList.Limit, offset)

	if err != nil {
		c.log.Error("error on GetCategories", logger.Error(err))
		return nil, err

	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&category.CategoryID,
			&category.Name,
			&category.CreatedAt,
		)
		if err != nil {
			c.log.Error("error on GetCategories rows", logger.Error(err))
			return nil, err

		}

		resp.Categories = append(resp.Categories, &category)

	}

	query = `
		SELECT 
			count(*)
		FROM
			 categories
	`

	err = c.db.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		c.log.Error("error on GetCategories count", logger.Error(err))
		return nil, err

	}

	resp.Count=count
	return &resp, nil
}
func (c *ContentRepo) UpdateCategory(ctx context.Context, req modles.Category) error {

	return nil
}
func (c *ContentRepo) DeleteCategory(ctx context.Context, id string) error {

	return nil
}
