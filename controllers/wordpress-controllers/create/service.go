package createWordPress

import "whm-api/utils/db/stacks"

type Service interface {
	CreateWordPressService(input *InputCreateWordPress) (stacks.ResponseStack, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateWordPressService(input *InputCreateWordPress) (stacks.ResponseStack, string) {
	stack, err := s.repository.CreateWordPressRepository(input)

	return stack.Response(), err
}
