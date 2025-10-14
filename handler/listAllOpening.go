package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAllOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Create Opening",
	})
}
