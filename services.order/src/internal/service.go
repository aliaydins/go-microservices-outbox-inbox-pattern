package order

import "github.com/aliaydins/oipattern/services.order/src/entity"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateOrder(order *entity.Order) error {

	err := s.repo.CreateOrder(order)
	if err != nil {
		return ErrOrderNotCreated
	}

	outboxItem := &entity.Outbox{
		EventType:     "OrderCreated",
		OrderID:       order.ID,
		CustomerEmail: order.CustomerEmail,
		Name:          order.Name,
		Amount:        order.Amount,
	}

	err = s.repo.CreateOutbox(outboxItem)
	if err != nil {
		return ErrOutboxItemNotCreated
	}

	return nil
}

func (s *Service) GetList() ([]OutboxDTO, error) {
	outbox, err := s.repo.GetOutboxList()
	if err != nil {
		return nil, ErrOutboxGetList
	}

	oDto := make([]OutboxDTO, 0)
	for _, e := range outbox {
		oDto = append(oDto, OutboxMapper(&e))
	}

	return oDto, nil
}
