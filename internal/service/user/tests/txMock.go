package tests

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TxMock struct {
	pgxpool.Tx
}

func (t *TxMock) Commit(_ context.Context) error {
	return nil
}

func (t *TxMock) Rollback(_ context.Context) error {
	return nil
}
