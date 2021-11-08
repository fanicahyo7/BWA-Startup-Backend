package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetCampaignTransaction(input GetCampaignTransactionInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaignTransaction(input GetCampaignTransactionInput) ([]Transaction, error) {
	transaction, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
