package postgres

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"
)

type ContentRepoI interface{

	CreateCategory(ctx context.Context,req *modles.Category)(*modles.Category,error)
	GetCategory(ctx context.Context,id string)(*modles.Category,error)
	GetCategories(ctx context.Context,getList modles.GetList)(*modles.GetCategories,error)
	UpdateCategory(ctx context.Context,req modles.Category)(error)
	DeleteCategory(ctx context.Context,id string)(error)
	
	CreateSubCategory(ctx context.Context,req modles.SubCategoryReq)(*modles.SubCategory,error)
	GetSubCategory(ctx context.Context,id string)(*modles.SubCategory,error)
	GetSubCategories(ctx context.Context,getList modles.GetList)(*modles.GetSubCategoriesLidtResp,error)
	UpdateSubCategory(ctx context.Context,req modles.SubCategoryReq)(error)
	DeleteSubCategory(ctx context.Context,id string)(error)

	CreateArticle(ctx context.Context,req *modles.ArticleReq)(*modles.ArticleResp,error)
	GetArticle(ctx context.Context,id string)(*modles.ArticleResp,error)
	GetArticles(ctx context.Context,limit,page int32)(*modles.GetArticleListResp,error)
	UpdateArticle(ctx context.Context,req modles.ArticleReq)(error)
	DeleteArticle(ctx context.Context,id string)(error)

}

type OwnerRepoI interface{
	Login(ctx context.Context,login *modles.LoginOwn)(*modles.Owner,error)
}

type CommonRepoI interface{
	CheckExists(ctx context.Context,req *modles.Common)(bool,error)
}

type ViwerRepoI interface{

	CreateViwer(ctx context.Context,req modles.Viewer)(*modles.Clamis,error)
	LogInViwer(ctx context.Context,req modles.LoginViwer)(*modles.Clamis,error)
	GetViwer(ctx context.Context,id string)(*modles.Viewer,error)
	GetViwers(ctx context.Context,list modles.GetList)(*modles.GetViwersResp,error)
	UpdateViwer(ctx context.Context,req modles.Viewer)(error)
	DeleteViwer(ctx context.Context,id string)(error)
	GetViwerClaims(ctx context.Context, id string) (*modles.Clamis, error)

	CreateComment(ctx context.Context,req *modles.CreateCommentReq)(*modles.Comment,error)
	GetComment(ctx context.Context,id string)(*modles.Comment,error)
	GetComments(ctx context.Context,getList modles.GetList)(*modles.GetCommments,error)
	UpdateComment(ctx context.Context,req modles.CreateCommentReq)(*modles.Comment,error)
	DeleteComment(ctx context.Context,id string)(error)
}