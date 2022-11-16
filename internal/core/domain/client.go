package domain

import "errors"

var errInvalidBalance = errors.New("wrong balance")

type Client struct {
	Id      int     `db:"id"`
	Balance float64 `db:"balance"`
}

func (c *Client) ValidateBalance() error {
	if c.Balance < 0 {
		return errInvalidBalance
	}

	return nil
}
