package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to-do-checklist/internal/domain"
)

func (h *Handler) signUp(ctx *gin.Context) {
	var input domain.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(ctx *gin.Context) {

}
