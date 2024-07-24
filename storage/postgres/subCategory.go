package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/google/uuid"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (c *ContentRepo) CreateSubCategory(ctx context.Context, req modles.SubCategoryReq) (*modles.SubCategory, error) {

	req.SubCategoryID = uuid.New()

	query := `	
		INSERT INTO
			sub_categories(
				sub_category_id,
				name,
				category_id
			)VALUES(
				$1,$2,$3
			)`

	_, err := c.db.Exec(
		ctx,
		query,
		req.SubCategoryID,
		req.Name,
		req.CategoryID,
	)

	if err != nil {
		c.log.Error("error on CreatsubCeategory ", logger.Error(err))
		return nil, err

	}
	resp, err := c.GetSubCategory(ctx, req.SubCategoryID.String())

	if err != nil {
		c.log.Error("error on  GetSubCategory", logger.Error(err))
		return nil, err

	}

	return resp, nil
}
func (c *ContentRepo) GetSubCategory(ctx context.Context, id string) (*modles.SubCategory, error) {
	var resp modles.SubCategory
	query := `
		SELECT
  		  *
		FROM 
   			sub_categories 
		WHERE 
  			sub_category_id = $1;

	`

	err := c.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&resp.SubCategoryID,
		&resp.Name,
		&resp.CreatedAt,
		&resp.CategoryID,
	)

	if err != nil {
		c.log.Error("error on Getsub_categorie", logger.Error(err))
		return nil, err

	}

	c.log.Debug("sucssesfully Get sub_categorie")
	return &resp, nil
}
func (c *ContentRepo) GetSubCategories(ctx context.Context, getList modles.GetList) (*modles.GetSubCategoriesLidtResp, error) {
	var subCategory modles.SubCategory
	var resp modles.GetSubCategoriesLidtResp
	var count int32

	offset := (getList.Pge - 1) * getList.Limit

	query := `
		SELECT 
			*
		FROM
			sub_categories
		LIMIT $1 OFFSET $2`

	rows, err := c.db.Query(ctx, query, getList.Limit, offset)

	if err != nil {
		c.log.Error("error on Get sub_categories", logger.Error(err))
		return nil, err

	}

	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&subCategory.SubCategoryID,
			&subCategory.Name,
			&subCategory.CreatedAt,
			&subCategory.CategoryID,
		)
		if err != nil {
			c.log.Error("error on GetSubCategories rows", logger.Error(err))
			return nil, err

		}

		resp.SubCategory = append(resp.SubCategory, &subCategory)

	}

	query = `
		SELECT 
			count(*)
		FROM
			sub_categories
	`

	err = c.db.QueryRow(ctx, query).Scan(&count)

	if err != nil {
		c.log.Error("error on SubCategory count", logger.Error(err))
		return nil, err

	}

	resp.Count = count

	return &resp, nil
}
func (c *ContentRepo) UpdateSubCategory(ctx context.Context, req modles.SubCategoryReq) error {
	return nil
}
func (c *ContentRepo) DeleteSubCategory(ctx context.Context, id string) error {
	return nil
}
