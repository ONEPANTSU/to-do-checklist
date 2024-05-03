package service

import (
	"errors"
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

func (s *TodoListService) GetListById(listID, userID int) (*domain.TodoList, error) {
	return s.repo.GetListByID(listID, userID)
}

func (s *TodoListService) UpdateList(list *domain.UpdateTodoList, listID, userID int) error {
	_, err := s.GetListById(listID, userID)
	if err != nil {
		return errors.New("list does not exist")
	}
	if err := list.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateList(list, listID, userID)
}

func (s *TodoListService) DeleteList(listID int, userID int) error {
	_, err := s.GetListById(listID, userID)
	if err != nil {
		return errors.New("list does not exist")
	}
	return s.repo.DeleteList(listID, userID)
}
