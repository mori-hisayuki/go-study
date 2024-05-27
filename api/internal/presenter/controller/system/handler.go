package system

import "github.com/gin-gonic/gin"

type SystemHandler struct{}


func(h *SystemHandler) Health(ctx *gin.Context){
	ctx.JSON(200, gin.H{
		"Status": "OK",
	})
}

func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}