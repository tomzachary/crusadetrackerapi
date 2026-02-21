package armies

import (
	"slices"
	"sync"
)

type Service struct {
	mu     sync.RWMutex
	armies []Army
}

func (s *Service) NewService() *Service {
	return &Service{armies: make([]Army, 0)}
}

func (s *Service) GetAllArmies() ([]Army, error) {
	return s.armies, nil
}

func (s *Service) CreateOrUpdateArmy(army Army) (Army, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.armies = append(s.armies, army)
	return army, nil
}

func (s *Service) DeleteArmy(id int) error {
	s.armies = slices.DeleteFunc(s.armies, func(a Army) bool {
		if a.Id == id {
			return true
		} else {
			return false
		}
	})
	return nil
}
