package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json: "transfer`
	FromAccount Account  `json: "from_account`
	ToAccount   Account  `json: "to_account`
	FromEntry   Entry    `json: "from_entry`
	ToEntry     Entry    `json: "to_entry`
}

func (s *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
				Amount: -arg.Amount,
				ID:     arg.FromAccountID,
			})
			if err != nil {
				return err
			}
			result.ToAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
				Amount: arg.Amount,
				ID:     arg.ToAccountID,
			})
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
				Amount: arg.Amount,
				ID:     arg.ToAccountID,
			})
			if err != nil {
				return err
			}
			result.FromAccount, err = q.AddAccountBalance(context.Background(), AddAccountBalanceParams{
				Amount: -arg.Amount,
				ID:     arg.FromAccountID,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}
