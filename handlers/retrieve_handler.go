package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kevalsabhani/go-cas-storage/cas"
)

type RetrieveContentResponse struct {
	Content string `json:"content"`
}

type Uri struct {
	Key string `uri:"key"`
}

func RetrieveContent(ctx *gin.Context) {
	uri := Uri{}
	if err := ctx.ShouldBindUri(&uri); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Print("============")
	fmt.Print(uri.Key)
	content, err := cas.RetrieveFromCAS(uri.Key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := RetrieveContentResponse{
		Content: content,
	}
	ctx.JSON(http.StatusOK, res)
}
