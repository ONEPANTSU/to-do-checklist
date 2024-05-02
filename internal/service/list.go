package service

import (
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) CreateList(list *domain.TodoList, userID int) (int, error) {
	return s.repo.CreateList(list, userID)
}

func (s *TodoListService) GetAllLists() *[]domain.TodoList {
	return s.repo.GetAllLists()
}

func (s *TodoListService) GetUsersLists(userID int) (*[]domain.TodoList, error) {
	return s.repo.GetUsersLists(userID)
}

func (s *TodoListService) GetListById(listID int, userID int) (*domain.TodoList, error) {
	return s.repo.GetListByID(listID, userID)
}
