package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"
)

type ContentRepoI interface{

	Createategory(ctx context.Context,req *modles.Category)(*modles.CategoryResp,error)
	GetCategory(ctx context.Context,id string)(*modles.CategoryResp,error)
	GetCategories(ctx context.Context,limit int32,page int32)(*modles.GetCategories,error)
	UpdateCategory(ctx context.Context,req modles.Category)(error)
	DeleteCategory(ctx context.Context,id string)(error)
	
	SubCreateCategory(ctx context.Context,req modles.SubCategory)(*modles.SubCategoryResp,error)
	GetSubCategory(ctx context.Context,id string)(*modles.SubCategoryResp,error)
	GetSubCategories(ctx context.Context,list modles.GetList)(*modles.SubCategoriesResp,error)
	UpdateSubCategory(ctx context.Context,req modles.SubCategory)(error)
	DeleteSubCategory(ctx context.Context,id string)(error)

	CreateViwer(ctx context.Context,req modles.ViewerReqReg)(*modles.Viewer,error)
	GetViwer(ctx context.Context,id string)(*modles.Viewer,error)
	GetViwers(ctx context.Context,list modles.GetList)(*modles.GetViwersResp,error)
	UpdateViwer(ctx context.Context,req modles.Viewer)(error)
	DeleteViwer(ctx context.Context,id string)(error)

}

type OwnerRepoI interface{
	Login(ctx context.Context,login *modles.LoginOwn)(*modles.Owner,error)
}

type CommonRepoI interface{
	CheckExists(ctx context.Context,req *modles.Common)(bool,error)
}

