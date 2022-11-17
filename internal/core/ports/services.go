package ports

import "github.com/krotovkk/bwgtransaction/internal/core/domain"

type ClientService interface {
	GetBalance(clientId int) (float64, error)
	ChangeBalance(clientId int, diff float64) (*domain.Client, error)
	Create() (*domain.Client, error)
}
