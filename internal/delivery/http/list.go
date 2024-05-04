package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"to-do-checklist/internal/domain"
)

// createList godoc
// @Security API Auth Key
// @Summary      Create List
// @Description  The route for creating new TODO List
// @Tags         TODO List
// @Accept       json
// @Produce      json
// @Param input body domain.TodoList true "TODO List info"
// @Success      200  {json}  	{"list_id": 1}
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/lists [post]
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

// getUsersLists godoc
// @Security API Auth Key
// @Summary      Get User's Lists
// @Description  The route for getting user's TODO Lists
// @Tags         TODO List
// @Accept       json
// @Produce      json
// @Success      200  {json}  	TodoListData
// @Failure      500  {object}  errorResponse
// @Router       /api/lists [get]
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

// getList godoc
// @Security API Auth Key
// @Summary      Get List by ID
// @Description  The route for getting TODO List by ID
// @Tags         TODO List
// @Accept       json
// @Produce      json
// @Param list_id path int true "TODO List ID"
// @Success      200  {json}  	domain.TodoList
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/lists/{list_id} [get]
func (h *Handler) getList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	listID, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	list, err := h.service.TodoList.GetListById(listID, userID)
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, list)
}

// updateList godoc
// @Security API Auth Key
// @Summary      Update List
// @Description  The route for editing TODO List
// @Tags         TODO List
// @Accept       json
// @Produce      json
// @Param list_id path int true "TODO List ID"
// @Param input body domain.UpdateTodoList true "Updating Fields"
// @Success      200  {string}  "string"
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/lists/{list_id} [put]
func (h *Handler) updateList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	listID, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	var input domain.UpdateTodoList
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.service.TodoList.UpdateList(&input, listID, userID); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("updated list %d", listID))
}

// deleteList godoc
// @Security API Auth Key
// @Summary      Delete List by ID
// @Description  The route for deleting TODO List
// @Tags         TODO List
// @Accept       json
// @Produce      json
// @Param list_id path int true "TODO List ID"
// @Success      200  {string}  "string"
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /api/lists/{list_id} [delete]
func (h *Handler) deleteList(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		return
	}
	listID, err := strconv.Atoi(ctx.Param("list_id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.service.TodoList.DeleteList(listID, userID); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, fmt.Sprintf("list deleted: %d", listID))
}
