package v1

import (
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/gin-gonic/gin"
)

func(h *handlers)CheckUser(ctx *gin.Context){
	
	var reqBody  modles.CheackViwer

	gin.Bind(&reqBody)

	isExists, err:= h.storage.GetCommonRepo().CheckExists(ctx,&modles.Common{
		TableName: "viwer",
		ColumnName: "gamil",
		ExpValue: reqBody.Gmail,
	})

	if err!=nil {

		ctx.JSON(200,err)
		return
	}

	if isExists{
		ctx.JSON(201,modles.CheckExists{
				IsExists: isExists,
				Status: "log-in",
			})
	}
	return
	 
		
	
}