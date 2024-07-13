package controllers

import (
	"strconv"
	"time"

	"github.com/OhVIton/ninsho-kun-backend/controllers/services"
	"github.com/gin-gonic/gin"
)

type AuthCodeController struct {
	s services.AuthCodeServicer
}

func (c *AuthCodeController) FetchAuthCode(ctx *gin.Context) {

	var afterSec int64
	afterSec, err := strconv.ParseInt(ctx.DefaultQuery("after", "0"), 10, 64)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "Invalid query parameter",
		})
		return
	}
	after := time.Unix(int64(afterSec), 0)

	code, err := c.s.FetchAuthCodeService(after)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if code == "" {
		ctx.JSON(404, gin.H{
			"error": "Auth code not found",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"authCode": code,
	})
}

func NewAuthCodeController(s services.AuthCodeServicer) *AuthCodeController {
	return &AuthCodeController{s}
}
