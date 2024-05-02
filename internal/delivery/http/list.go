package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"to-do-checklist/internal/domain"
)

func (h *Handler) createList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	var input domain.TodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	listID, err := h.service.TodoList.CreateList(&input, userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"list_id": listID,
	})
}

func (h *Handler) getUsersLists(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	lists, err := h.service.TodoList.GetUsersLists(userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, TodoListsData{
		Data: *lists,
	})
}
func (h *Handler) getList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	listID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
	}
	list, err := h.service.TodoList.GetListById(listID, userID)
	if err != nil {
		return
	}
	ctx.JSON(http.StatusOK, list)

}
func (h *Handler) updateList(ctx *gin.Context) {}
func (h *Handler) deleteList(ctx *gin.Context) {}
