package db

//import (
//	"context"
//	"database/sql"
//	"fmt"
//)
//
//type Store struct {
//	*Queries
//	db *sql.DB
//}
//
//func NewStore(db *sql.DB) *Store {
//	return &Store{
//		Queries: New(db),
//		db:      db,
//	}
//}
//
//func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
//	tx, err := s.db.BeginTx(ctx, nil)
//	if err != nil {
//		return err
//	}
//
//	q := New(tx)
//	err = fn(q)
//	if err != nil {
//		if rbErr := tx.Rollback(); rbErr != nil {
//			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
//		}
//		return err
//	}
//	return tx.Commit()
//}
//
//type TransferTxParams struct {
//	FromAccountID int64 `json:"from_account_id"`
//	ToAccountID   int64 `json:"to_account_id"`
//	Account       int64 `json:"account"`
//}
//
//type TransferTxResult struct {
//	Transfer      Transfer `json:"transfer"`
//	FromAccountID int64    `json:"from_account_id"`
//	ToAccountID   int64    `json:"to_account_id"`
//	FromEntry     Entry    `json:"from_entry"`
//	ToEntry       Entry    `json:"to_entry"`
//}
//
//func (s Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
//	var result TransferTxResult
//
//	err := s.execTx(ctx, func(q *Queries) error {
//		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
//			FromAccountID: sql.NullInt64{},
//			ToAccountID:   sql.NullInt64{},
//			Amount:        0,
//		})
//		return nil
//	})
//
//	return result, err
//}
