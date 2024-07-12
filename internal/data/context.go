package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
)

type contextKey int
type contextTransactionIDKey string

const (
	TransactionKey   contextKey              = iota
	TransactionIDkey contextTransactionIDKey = "txid"
)

// NewContext returns a new context with transaction stored in it.
// Upon error, the original context is still returned along with an error
func (c *ConnectionFactory) NewContext(ctx context.Context) (context.Context, error) {
	tx, err := c.newTransaction()
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, TransactionIDkey, tx.txid)
	ctx = context.WithValue(ctx, TransactionKey, tx)

	return ctx, nil
}

// TxContext creates a new transaction context from context.Background()
func (c *ConnectionFactory) TxContext() (ctx context.Context, err error) {
	return c.NewContext(context.Background())
}

// Resolve resolves the current transaction according to the rollback flag.
func Resolve(ctx context.Context) error {
	tx, ok := ctx.Value(TransactionKey).(*txFactory)
	if !ok {
		return fmt.Errorf("could not retrieve transaction from context")
	}
	if tx.resolved {
		return nil
	}
	tx.resolved = true
	postCommitActions := tx.postCommitActions
	tx.postCommitActions = nil
	if tx.markedForRollback() {
		if err := tx.tx.Rollback(); err != nil {
			return fmt.Errorf("could not rollback transaction: %v", err)
		}
		log.Infof("Rolled back transaction")
	} else {
		if err := tx.tx.Commit(); err != nil {
			// TODO:  what does the user see when this occurs? seems like they will get a false positive
			return fmt.Errorf("could not commit transaction: %v", err)
		}
		for _, f := range postCommitActions {
			f()
		}
	}
	return nil
}

func Begin(ctx context.Context) error {
	tx, ok := ctx.Value(TransactionKey).(*txFactory)
	if !ok {
		return fmt.Errorf("could not retrieve transaction from context")
	}

	err := tx.begin()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %v", err)
	}
	return nil
}

func AddPostCommitAction(ctx context.Context, f func()) error {
	tx, ok := ctx.Value(TransactionKey).(*txFactory)
	if !ok {
		return fmt.Errorf("could not retrieve transaction from context")
	}

	tx.postCommitActions = append(tx.postCommitActions, f)
	return nil
}

// FromContext Retrieves the transaction from the context.
func FromContext(ctx context.Context) (*sql.Tx, error) {
	transaction, ok := ctx.Value(TransactionKey).(*txFactory)
	if !ok {
		return nil, fmt.Errorf("could not retrieve transaction from context")
	}
	return transaction.tx, nil
}

// MarkForRollback flags the transaction stored in the context for rollback and logs whatever error caused the rollback
func MarkForRollback(ctx context.Context, err error) {
	transaction, ok := ctx.Value(TransactionKey).(*txFactory)
	if !ok {
		log.Errorf("failed to mark transaction for rollback: could not retrieve transaction from context")
		return
	}
	log.Infof("Marked transaction for rollback, err: %v", err)
	transaction.rollbackFlag = true
}
