package listStacks

import (
	"whm-api/utils/db/stacks"
)

type Service interface {
	ListStacksService() ([]stacks.ResponseStack, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) ListStacksService() ([]stacks.ResponseStack, string) {
	list, err := s.repository.ListStacksRepository()

	res := []stacks.ResponseStack{}

	for _, stack := range list {
		res = append(res, stack.Response())
	}

	return res, err
}
