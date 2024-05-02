package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const authHeader = "Authorization"
const userContext = "userID"

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}
	userID, err := h.service.Authorization.ValidateToken(headerParts[1])
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}
	ctx.Set(userContext, userID)
}

func getUserID(ctx *gin.Context) (int, error) {
	userID, ok := ctx.Get(userContext)
	if !ok {
		newErrorResponse(ctx, http.StatusBadRequest, "user id not found")
		return 0, errors.New("user id not found")
	}
	id, ok := userID.(int)
	if !ok {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid user id")
		return 0, errors.New("invalid user id")
	}
	return id, nil
}
