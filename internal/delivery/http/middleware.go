package http

import (
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
