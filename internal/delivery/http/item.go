package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"to-do-checklist/internal/domain"
)

// createItem godoc
// @Security API Auth Key
// @Summary      Create Item
// @Description  The route for creating new TODO Item in TODO List
// @Tags         TODO Item
// @Accept       json
// @Produce      json
// @Param input body domain.TodoItem true "TODO Item info"
// @Success      200  {json}  	{"item_id": 1}
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/items [post]
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

// getItems godoc
// @Security API Auth Key
// @Summary      Get Items from List
// @Description  The route for getting all TODO Items from TODO List
// @Tags         TODO Item
// @Accept       json
// @Produce      json
// @Param list_id path int true "TODO List ID"
// @Success      200  {json}  	TodoItemData
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/lists/{list_id}/items [get]
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

// getItem godoc
// @Security API Auth Key
// @Summary      Get Item by ID
// @Description  The route for getting TODO Item by ID
// @Tags         TODO Item
// @Accept       json
// @Produce      json
// @Param item_id path int true "TODO Item ID"
// @Success      200  {json}  	domain.TodoItem
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/items/{item_id} [get]
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

// updateItem godoc
// @Security API Auth Key
// @Summary      Update Item
// @Description  The route for editing TODO Item
// @Tags         TODO Item
// @Accept       json
// @Produce      json
// @Param item_id path int true "TODO Item ID"
// @Param input body domain.UpdateTodoItem true "Updating Fields"
// @Success      200  {string}  "string"
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/items/{item_id} [put]
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

// deleteItem godoc
// @Security API Auth Key
// @Summary      Delete Item by ID
// @Description  The route for deleting TODO Item
// @Tags         TODO Item
// @Accept       json
// @Produce      json
// @Param item_id path int true "TODO Item ID"
// @Success      200  {string}  "string"
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/items/{item_id} [delete]
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
