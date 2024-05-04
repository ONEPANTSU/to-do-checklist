package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"to-do-checklist/internal/domain"
)

func (h *Handler) createItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	var input domain.TodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	itemID, err := h.service.TodoItem.CreateItem(&input, userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{"item_id": itemID})
}
func (h *Handler) getItems(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	listID, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	items, err := h.service.TodoItem.GetItems(listID, userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, TodoItemData{*items})
}

func (h *Handler) getItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	item, err := h.service.TodoItem.GetItemById(itemID, userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.UpdateTodoItem
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.TodoItem.UpdateItem(&input, itemID, userID); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("updated item %d", itemID))
}

func (h *Handler) deleteItem(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	itemID, err := strconv.Atoi(ctx.Param("item_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.TodoItem.DeleteItem(itemID, userID); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("deleted item %d", itemID))
}
