package service

import (
	"errors"
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

type TodoListService struct {
	listRepo repository.TodoList
}

func newTodoListService(listRepo repository.TodoList) *TodoListService {
	return &TodoListService{listRepo: listRepo}
}

func (s *TodoListService) CreateList(list *domain.TodoList, userID int) (int, error) {
	return s.listRepo.CreateList(list, userID)
}

func (s *TodoListService) GetAllLists() *[]domain.TodoList {
	return s.listRepo.GetAllLists()
}

func (s *TodoListService) GetUsersLists(userID int) (*[]domain.TodoList, error) {
	return s.listRepo.GetUsersLists(userID)
}

func (s *TodoListService) GetListById(listID, userID int) (*domain.TodoList, error) {
	return s.listRepo.GetListByID(listID, userID)
}

func (s *TodoListService) UpdateList(list *domain.UpdateTodoList, listID, userID int) error {
	_, err := s.GetListById(listID, userID)
	if err != nil {
		return errors.New("list does not exist")
	}
	if err := list.Validate(); err != nil {
		return err
	}
	return s.listRepo.UpdateList(list, listID, userID)
}

func (s *TodoListService) DeleteList(listID int, userID int) error {
	_, err := s.GetListById(listID, userID)
	if err != nil {
		return errors.New("list does not exist")
	}
	return s.listRepo.DeleteList(listID, userID)
}
