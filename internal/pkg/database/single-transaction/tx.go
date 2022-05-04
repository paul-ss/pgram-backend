package single_transaction

import (
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type tx struct {
	id   string
	conn *conn
}

func (tx *tx) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, ErrNotImplemented
}

func (tx *tx) BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error) {
	return ErrNotImplemented
}

func (tx *tx) Commit(ctx context.Context) error {
	if tx.conn.savePoint == nil {
		return nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return err
	}

	_, err = connTx.Exec(ctx, tx.conn.savePoint.Release(tx.id))
	return err
}

func (tx *tx) Rollback(ctx context.Context) error {
	if tx.conn.savePoint == nil {
		return nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return err
	}

	_, err = connTx.Exec(ctx, tx.conn.savePoint.Rollback(tx.id))
	return err
}

func (tx *tx) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	return 0, ErrNotImplemented
}

func (tx *tx) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}

func (tx *tx) LargeObjects() pgx.LargeObjects {
	return pgx.LargeObjects{}
}

func (tx *tx) Prepare(ctx context.Context, name string, sql string) (*pgconn.StatementDescription, error) {
	return nil, ErrNotImplemented
}

func (tx *tx) Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error) {
	if tx.conn.savePoint == nil {
		return nil, nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return nil, err
	}

	return connTx.Exec(ctx, sql, arguments...)
}

func (tx *tx) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	if tx.conn.savePoint == nil {
		return nil, nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return nil, err
	}

	return connTx.Query(ctx, sql, args...)
}

func (tx *tx) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	if tx.conn.savePoint == nil {
		return nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return nil
	}

	return connTx.QueryRow(ctx, sql, args...)
}

func (tx *tx) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	if tx.conn.savePoint == nil {
		return nil, nil // save point is not supported
	}

	tx.conn.Lock()
	defer tx.conn.Unlock()

	connTx, err := tx.conn.beginOnce()
	if err != nil {
		return nil, err
	}

	return connTx.QueryFunc(ctx, sql, args, scans, f)
}

func (tx *tx) Conn() *pgx.Conn {
	return nil
}
