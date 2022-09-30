package notification

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetList() ([]InboxDTO, error) {
	outbox, err := s.repo.GetInboxList()
	if err != nil {
		return nil, err
	}

	iDto := make([]InboxDTO, 0)
	for _, e := range outbox {
		iDto = append(iDto, InboxMapper(&e))
	}

	return iDto, nil
}
