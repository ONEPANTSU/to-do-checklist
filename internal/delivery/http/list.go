package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createList(ctx *gin.Context) {}
func (h *Handler) getAllLists(ctx *gin.Context) {
	id, _ := ctx.Get(userContext)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) getList(ctx *gin.Context)    {}
func (h *Handler) updateList(ctx *gin.Context) {}
func (h *Handler) deleteList(ctx *gin.Context) {}
