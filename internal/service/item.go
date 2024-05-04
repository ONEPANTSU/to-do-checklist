package service

import (
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	listRepo repository.TodoList
}

func newTodoItemService(iterRepo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{itemRepo: iterRepo, listRepo: listRepo}
}

func (s *TodoItemService) CreateItem(item *domain.TodoItem, userID int) (int, error) {
	if _, err := s.listRepo.GetListByID(item.ListID, userID); err != nil {
		return 0, err
	}
	return s.itemRepo.CreateItem(item)
}

func (s *TodoItemService) GetItems(listID, userID int) (*[]domain.TodoItem, error) {
	if _, err := s.listRepo.GetListByID(listID, userID); err != nil {
		return nil, err
	}
	return s.itemRepo.GetItems(listID)
}

func (s *TodoItemService) GetItemById(itemID, userID int) (*domain.TodoItem, error) {
	return s.itemRepo.GetItemByID(itemID, userID)
}

func (s *TodoItemService) UpdateItem(item *domain.UpdateTodoItem, itemID, userID int) error {
	if _, err := s.itemRepo.GetItemByID(itemID, userID); err != nil {
		return err
	}
	if err := item.Validate(); err != nil {
		return err
	}
	return s.itemRepo.UpdateItem(item, itemID)
}

func (s *TodoItemService) DeleteItem(itemID, userID int) error {
	if _, err := s.itemRepo.GetItemByID(itemID, userID); err != nil {
		return err
	}
	return s.itemRepo.DeleteItem(itemID)
}
