package clientservice

import (
	"github.com/krotovkk/bwgtransaction/internal/core/domain"
	"github.com/krotovkk/bwgtransaction/internal/core/ports"
)

type Options struct {
	ClientRepo  ports.ClientRepository
	HistoryRepo ports.HistoryRepository
}

type ClientService struct {
	clientRepo  ports.ClientRepository
	historyRepo ports.HistoryRepository
}

func NewClientService(options *Options) *ClientService {
	return &ClientService{
		clientRepo:  options.ClientRepo,
		historyRepo: options.HistoryRepo,
	}
}

func (cs *ClientService) GetBalance(clientId int) (float64, error) {
	client, err := cs.clientRepo.GetClient(clientId)

	if err != nil {
		return 0, err
	}

	return client.Balance, nil
}

func (cs *ClientService) ChangeBalance(clientId int, diff float64) (*domain.Client, error) {
	client, err := cs.clientRepo.GetClient(clientId)

	if err != nil {
		return nil, err
	}

	client.Balance += diff

	err = client.ValidateBalance()

	if err != nil {
		return nil, err
	}

	client, err = cs.clientRepo.UpdateBalance(client)

	if err != nil {
		return nil, err
	}

	_ = cs.historyRepo.Save(client)

	return client, nil
}
