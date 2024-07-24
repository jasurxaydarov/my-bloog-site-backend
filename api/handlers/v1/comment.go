package v1

import (
	"context"
	"jasurxaydarov/my-bloog-site-backend/modles"
	"jasurxaydarov/my-bloog-site-backend/pgx/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// SubCategory

func (h *handlers) AddComment(ctx *gin.Context) {
	var (
		req  modles.CreateCommentReq
		reqq = &modles.CreateCommentReq{}
	)
	claim := Auth(ctx)

	if claim == nil {
		ctx.JSON(401, "Unauth")
	}
	err := ctx.Bind(&req)

	if err != nil {
		ctx.JSON(401, err)
		return
	}
	helpers.DataParser(req, reqq)
	reqq.ViewerID, _ = uuid.Parse(claim.UserId)
	resp, err := h.storage.GetViwerRepo().CreateComment(context.Background(), reqq)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *handlers) GetComment(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := h.storage.GetViwerRepo().GetComment(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)

}

func (h *handlers) GetComments(ctx *gin.Context) {

	var list modles.GetList

	limit := ctx.Query("limit")
	page := ctx.Query("page")

	list.Limit = helpers.GetLimit(limit)
	list.Pge = helpers.GetPage(page)

	resp, err := h.storage.GetViwerRepo().GetComments(context.Background(), list)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}

func (h *handlers) UpdateComment(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := h.storage.GetContentRepo().GetCategory(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)

}

func (h *handlers) DeleteComment(ctx *gin.Context) {

	id := ctx.Param("id")

	resp, err := h.storage.GetContentRepo().GetCategory(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, resp)
}
