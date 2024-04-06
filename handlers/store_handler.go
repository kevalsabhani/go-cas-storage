package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevalsabhani/go-cas-storage/cas"
)

type StoreContentRequest struct {
	Content string `json:"content"`
}

type StoreContentResponse struct {
	Key string `json:"key"`
}

func StoreContent(ctx *gin.Context) {

	req := &StoreContentRequest{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	key, err := cas.StoreInCAS(req.Content)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := StoreContentResponse{
		Key: key,
	}

	ctx.JSON(http.StatusCreated, res)
}
