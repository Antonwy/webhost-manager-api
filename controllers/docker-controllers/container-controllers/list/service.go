package listContainers

import "github.com/docker/docker/api/types"

type Service interface {
	ListContainersService() ([]types.Container, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListContainersService() ([]types.Container, string) {
	resultCreateStudent, errCreateStudent := s.repository.ListContainersRepository()

	return resultCreateStudent, errCreateStudent
}
