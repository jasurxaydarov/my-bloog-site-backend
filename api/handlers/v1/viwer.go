package v1

import (
	"encoding/json"
	"jasurxaydarov/my-bloog-site-backend/mail"
	"jasurxaydarov/my-bloog-site-backend/modles"

	"github.com/gin-gonic/gin"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *handlers) CheckUser(ctx *gin.Context) {

	var reqBody modles.CheackViwer

	err := ctx.Bind(&reqBody)

	if err != nil {

		ctx.JSON(500, err)
		return
	}

	isExists, err := h.storage.GetCommonRepo().CheckExists(ctx, &modles.Common{
		TableName:  "viwers",
		ColumnName: "gmail",
		ExpValue:   reqBody.Gmail,
	})

	if err != nil {

		ctx.JSON(500, err)
		return
	}

	if isExists {
		ctx.JSON(201, modles.CheckExists{
			IsExists: isExists,
			Status:   "log-in",
		})
		return
	}

	otp := modles.OtpData{
		Otp:   mail.GenerateOtp(6),
		Gmail: reqBody.Gmail,
	}

	otpdataB, err := json.Marshal(otp)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	err = h.cache.Set(ctx, reqBody.Gmail, string(otpdataB), 60)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	err = mail.SendMail([]string{reqBody.Gmail}, otp.Otp)

	if err != nil {
		h.log.Error("errrr on Send mail",logger.Error(err))
		ctx.JSON(500, err)
		return
	}
	
		ctx.JSON(201, modles.CheckExists{
			IsExists: isExists,
			Status:   "registr",
		})
	ctx.JSON(201, "sent")
}
