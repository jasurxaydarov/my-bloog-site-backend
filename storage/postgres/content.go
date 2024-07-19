package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/saidamir98/udevs_pkg/logger"
)

type ContentRepo struct {
	db  *pgx.Conn
	log logger.LoggerI
}

func NewContent(db *pgx.Conn, log logger.LoggerI) ContentRepoI {

	return &ContentRepo{db, log}
}

func (c *ContentRepo) Createategory(ctx context.Context, req *modles.Category) (*modles.CategoryResp, error) {

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
func (c *ContentRepo) GetCategory(ctx context.Context, id string) (*modles.CategoryResp, error) {

	var resp modles.CategoryResp
	query := `
		SELECT
  		  name 
		FROM 
   			categories 
		WHERE 
  			category_id = $1;

	`

	err := c.db.QueryRow(ctx, query, id).Scan(&resp.Name)

	if err != nil {
		c.log.Error("error on GetCategory", logger.Error(err))
		return nil, err

	}

	c.log.Debug("sucssesfully Get category")
	return &resp, nil
}
func (c *ContentRepo) GetCategories(ctx context.Context, limit int32, page int32) (*modles.GetCategories, error) {
	var category modles.Category
	var resp []*modles.Category
	var count int32

	offset := (page - 1) * limit

	query := `
		SELECT 
			*
		FROM
			categories
		LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(ctx, query, limit, offset)

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

		resp = append(resp, &category)

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

	return &modles.GetCategories{Categories: resp, Count: count}, nil
}
func (c *ContentRepo) UpdateCategory(ctx context.Context, req modles.Category) error {

	return nil
}
func (c *ContentRepo) DeleteCategory(ctx context.Context, id string) error {

	return nil
}

func (c *ContentRepo)SubCreateCategory(ctx context.Context,req modles.SubCategory)(*modles.SubCategoryResp,error){
	return nil,nil
}
func (c *ContentRepo)GetSubCategory(ctx context.Context,id string)(*modles.SubCategoryResp,error){
	return nil,nil
}
func (c *ContentRepo)GetSubCategories(ctx context.Context,list modles.GetList)(*modles.SubCategoriesResp,error){
	return nil,nil
}
func (c *ContentRepo)UpdateSubCategory(ctx context.Context,req modles.SubCategory)(error){
	return nil
}
func (c *ContentRepo)DeleteSubCategory(ctx context.Context,id string)(error){
	return nil
}


/*
	INSERT INTO viwers (viwer_id, fullname, username, gmail, password)
VALUES
    ('123e4567-e89b-12d3-a456-426614174000', 'John Doe', 'johndoe', 'xaydarovjasur2005@gmail.com', 'password123'),
    ('123e4567-e89b-12d3-a456-426614174001', 'Jane Smith', 'janesmith', 'janesmith@gmail.com', 'password456'),
    ('123e4567-e89b-12d3-a456-426614174002', 'Alice Johnson', 'alicej', 'alicej@gmail.com', 'password789');

*/
