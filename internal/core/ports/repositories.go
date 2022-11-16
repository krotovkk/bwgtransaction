package ports

import "github.com/krotovkk/bwgtransaction/internal/core/domain"

type ClientRepository interface {
	GetClient(clientId int) (*domain.Client, error)
	UpdateBalance(client *domain.Client) (*domain.Client, error)
}

type HistoryRepository interface {
	Save(client *domain.Client) error
}
