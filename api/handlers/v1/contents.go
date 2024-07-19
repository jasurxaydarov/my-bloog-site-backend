package v1

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"
	"jasurxaydarov/my-bloog-site-backend/pgx/helpers"


	"github.com/gin-gonic/gin"
)


func (h *handlers)CreateCategory(ctx *gin.Context){

	var (
		req 	modles.Category
		reqq 	=&modles.Category{}
	)

	err:=ctx.Bind(&req)

	if err!= nil{
		ctx.JSON(401,err)
		return
	}
	helpers.DataParser(req,reqq)

	resp, err:=h.storage.GetContentRepo().Createategory(context.Background(),reqq)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(201,resp)


}

func (h *handlers)GetCategory(ctx *gin.Context){


	id:=ctx.Param("id")

	resp, err:= h.storage.GetContentRepo().GetCategory(context.Background(),id)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)

}

func(h *handlers)GetCategories(ctx *gin.Context){

	var list modles.GetList

	limit:=ctx.Query("limit")
	page:=ctx.Query("page")

	list.Limit=helpers.GetLimit(limit)
	list.Pge=helpers.GetPage(page)

	resp,err:=h.storage.GetContentRepo().GetCategories(context.Background(),list.Limit,list.Pge)

	if err!= nil{
		ctx.JSON(500,err)
		return
	}

	ctx.JSON(200,resp)
}


