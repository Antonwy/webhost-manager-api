package removeStack

type Service interface {
	RemoveStackService(id string) string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) RemoveStackService(id string) string {
	return s.repository.RemoveStackRepository(id)
}
