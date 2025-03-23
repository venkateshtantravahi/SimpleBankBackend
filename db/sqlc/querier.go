// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accounts, error)
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entries, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfers, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	DeleteTransfer(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Accounts, error)
	GetEntry(ctx context.Context, id int64) (Entries, error)
	GetTransfer(ctx context.Context, id int64) (Transfers, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Accounts, error)
	ListEntriesByAccount(ctx context.Context, arg ListEntriesByAccountParams) ([]Entries, error)
	ListTransfersBetweenAccounts(ctx context.Context, arg ListTransfersBetweenAccountsParams) ([]Transfers, error)
	ListTransfersFromAccount(ctx context.Context, arg ListTransfersFromAccountParams) ([]Transfers, error)
	ListTransfersToAccount(ctx context.Context, arg ListTransfersToAccountParams) ([]Transfers, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accounts, error)
	UpdateEntry(ctx context.Context, arg UpdateEntryParams) (Entries, error)
	UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (Transfers, error)
}

var _ Querier = (*Queries)(nil)
