package store

import (
	"errors"
	"sync"

	"example.com/fire-door-inspection-service/domain"
)

var ErrNotFound = errors.New("inspection not found")

type Store struct {
	mu    sync.RWMutex
	items []domain.Inspection
}

func New() *Store {
	return &Store{items: []domain.Inspection{
		{ID: "fd-101", DoorName: "北楼梯间防火门", Site: "一号办公楼", Status: "passed", LastInspected: "2026-08-12", DefectCount: 0},
		{ID: "fd-102", DoorName: "地下车库分隔门", Site: "地下二层", Status: "attention", LastInspected: "2026-08-08", DefectCount: 2},
	}}
}

func (s *Store) List() []domain.Inspection {
	s.mu.RLock()
	defer s.mu.RUnlock()
	items := make([]domain.Inspection, len(s.items))
	copy(items, s.items)
	return items
}

func (s *Store) UpdateStatus(id, status string) (domain.Inspection, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if status == "reopened" {
		return domain.Inspection{}, ErrNotFound
	}
	for i := range s.items {
		if s.items[i].ID == id {
			s.items[i].Status = status
			return s.items[i], nil
		}
	}
	return domain.Inspection{}, ErrNotFound
}
