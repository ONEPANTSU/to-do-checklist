package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"to-do-checklist/internal/domain"
)

// signUp godoc
// @Summary      Sign up
// @Description  The route for registration
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input body domain.User true "Register info"
// @Success      200  {json}  	{"id": 1}
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /auth/sign-up [post]
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

// signIn godoc
// @Summary      Sign in
// @Description  The route for log-in
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param input body domain.SignIn true "Log-in info"
// @Success      200  {json}  	{"token": "string"}
// @Failure      400  {object}  errorResponse
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(ctx *gin.Context) {
	var input domain.SignIn

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Authorization.GenerateToken(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
