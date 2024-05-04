package http

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"to-do-checklist/internal/domain"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) errorResponse {
	logrus.Error(message)
	ctx.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
	return errorResponse{Message: message}
}

type TodoListsData struct {
	Data []domain.TodoList `json:"data"`
}

type TodoItemData struct {
	Data []domain.TodoItem `json:"data"`
}
