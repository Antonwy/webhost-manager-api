package createWordPress

type Service interface {
	CreateWordPressService(input *InputCreateWordPress) (*InputCreateWordPress, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateWordPressService(input *InputCreateWordPress) (*InputCreateWordPress, string) {
	wordPress, err := s.repository.CreateWordPressRepository(input)

	return wordPress, err
}
