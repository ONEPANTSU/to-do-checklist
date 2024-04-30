package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) ErrorResponse {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
	return ErrorResponse{Message: message}
}
