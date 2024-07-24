package v1

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"
	"jasurxaydarov/my-bloog-site-backend/pgx/helpers"

	"github.com/gin-gonic/gin"
)

// SubCategory



func (h *handlers)CreateSubCategory(ctx *gin.Context){
	var (
		req 	modles.SubCategoryReq
		reqq 	=&modles.SubCategoryReq{}
	)

	err:=ctx.Bind(&req)

	if err!= nil{
		ctx.JSON(401,err)
		return
	}
	helpers.DataParser(req,reqq)

	resp, err:=h.storage.GetContentRepo().CreateSubCategory(context.Background(),*reqq)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,resp)

}


func (h *handlers) GetSubCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := h.storage.GetContentRepo().GetCategory(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)

}

func (h *handlers) GetSubCategories(ctx *gin.Context) {

	var list modles.GetList

	limit := ctx.Query("limit")
	page := ctx.Query("page")

	list.Limit = helpers.GetLimit(limit)
	list.Pge = helpers.GetPage(page)

	resp, err := h.storage.GetContentRepo().GetSubCategories(context.Background(), list)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}


func (h *handlers) UpdateSubCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := h.storage.GetContentRepo().GetCategory(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)

}

func (h *handlers) DeleteSubCategory(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := h.storage.GetContentRepo().GetCategory(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}