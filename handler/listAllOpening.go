package handler

import (
	"net/http"

	"github.com/andamorim2206/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListAllOpeningHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil{
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
