package armies

import "sync"

type Service struct {
	mu     sync.RWMutex
	armies []Army
}

func (s *Service) NewService() *Service {
	return &Service{armies: make([]Army, 0)}
}

func (s *Service) GetAllArmies() []Army {
	return s.armies
}

func (s *Service) CreateOrUpdateArmy() {

}

func (s *Service) DeleteArmy() {

}
