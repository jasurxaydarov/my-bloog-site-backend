package v1

import (
	"context"
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

	err = h.cache.Set(ctx, reqBody.Gmail, string(otpdataB), 120)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	err = mail.SendMail([]string{reqBody.Gmail}, otp.Otp)

	if err != nil {
		h.log.Error("errrr on Send mail", logger.Error(err))
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, modles.CheckExists{
		IsExists: isExists,
		Status:   "register",
	})
	ctx.JSON(201, "sent")
}

func (h *handlers) CheckOtp(ctx *gin.Context) {
	var reqBody modles.OtpData

	err := ctx.Bind(&reqBody)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	data, err := h.cache.GetDell(ctx, reqBody.Gmail)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if data == "" {
		ctx.JSON(201, "otp is expired")
		return
	}

	var cacheData modles.OtpData

	json.Unmarshal([]byte(data), &cacheData)

	ctx.JSON(201, modles.CheckOtpRep{IsRight: cacheData.Otp == reqBody.Otp})
}

func (h *handlers) SignUp(ctx *gin.Context) {
	var reqBody modles.ViewerReqReg
	var otpData modles.OtpData

	err := ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		h.log.Error("errrr on ShouldBindJSON", logger.Error(err))
		return
	}

	otpSData, err := h.cache.GetDell(ctx, reqBody.Gmail)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if otpSData == "" {
		ctx.JSON(201, "otp is expired")
		return
	}

	err = json.Unmarshal([]byte(otpSData), &otpData)

	if otpData.Otp != reqBody.Otp {
		ctx.JSON(405, "incorrect otp")
		return
	}

	_, err = h.storage.GetContentRepo().CreateViwer(context.Background(), reqBody)

	if otpData.Otp != reqBody.Otp {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, "succesfully registretd")
}
