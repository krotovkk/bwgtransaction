package ports

import "github.com/krotovkk/bwgtransaction/internal/core/domain"

type ClientRepository interface {
	GetClient(clientId int) (*domain.Client, error)
	UpdateClient(client *domain.Client) (*domain.Client, error)
	CreateClient(c *domain.Client) (*domain.Client, error)
}

type HistoryRepository interface {
	Save(clientId int, diff float64) error
}
