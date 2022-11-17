package clientrepository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/krotovkk/bwgtransaction/internal/core/domain"
)

type ClientRepository struct {
	conn *pgx.Conn
}

func NewClientRepository(conn *pgx.Conn) *ClientRepository {
	return &ClientRepository{
		conn: conn,
	}
}

func (r *ClientRepository) GetClient(clientId int) (*domain.Client, error) {
	query, args, err := squirrel.Select("id", "balance").
		From("clients").
		Where(squirrel.Eq{"id": clientId}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("ClientRepository.GetClient: to sql: %w", err)
	}

	var client domain.Client
	err = r.conn.QueryRow(context.Background(), query, args...).Scan(&client.Id, &client.Balance)

	if err != nil {
		return nil, fmt.Errorf("ClientRepository.GetClient: select: %w", err)
	}

	return &client, nil
}

func (r *ClientRepository) UpdateClient(client *domain.Client) (*domain.Client, error) {
	query, args, err := squirrel.Update("clients").
		Set("balance", client.Balance).
		Where(squirrel.Eq{"id": client.Id}).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("ClientRepository.UpdateClient: to sql: %w", err)
	}

	_, err = r.conn.Exec(context.Background(), query, args...)
	if err != nil {
		return nil, fmt.Errorf("ClientRepository.UpdateClient: insert: %w", err)
	}

	return client, nil
}

func (r *ClientRepository) CreateClient(c *domain.Client) (*domain.Client, error) {
	query, args, err := squirrel.Insert("clients").
		Columns("balance").
		Values(c.Balance).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return nil, fmt.Errorf("ClientRepository.CreateClient: to sql: %w", err)
	}

	err = r.conn.QueryRow(context.Background(), query, args...).Scan(&c.Id)
	if err != nil {
		return nil, fmt.Errorf("ClientRepository.CreateClient: insert: %w", err)
	}

	return c, nil
}
