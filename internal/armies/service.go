package armies

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAllArmies() ([]Army, error) {
	return s.repo.GetAllArmies()
}

func (s *Service) CreateArmy(army Army) (Army, error) {
	return s.repo.CreateArmy(army)
}

func (s *Service) UpdateArmy(id int, army Army) (Army, error) {
	return s.repo.UpdateArmy(id, army)
}

func (s *Service) DeleteArmy(id int) error {
	return s.repo.DeleteArmy(id)
}
