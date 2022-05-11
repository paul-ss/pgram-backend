package single_transaction

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"sync"
)

const defaultDnsString = "user=tst password=123 host=localhost port=5430 database=tst sslmode=disable"

var ErrNotImplemented = errors.New("not implemented")

var drv *Driver

type Driver struct {
	sync.Mutex
	dbConn *pgx.Conn
	cn     *conn
}

func (d *Driver) deleteConn() {
	d.Lock()
	defer d.Unlock()
	d.cn = nil
}

func newDriver(dns string) (*Driver, error) {
	c, err := pgx.Connect(context.Background(), dns)

	if err != nil {
		return nil, err
	}

	if err := c.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Driver{
		dbConn: c,
	}, nil
}

func InitDriver(dns string) error {
	if dns == "" {
		dns = defaultDnsString
	}

	if drv == nil {
		var err error
		drv, err = newDriver(dns)

		if err != nil {
			drv = nil
			return err
		}
	}

	return nil
}

func TeardownDriver() error {
	return drv.dbConn.Close(context.Background())
}

type conn struct {
	sync.Mutex
	tx        pgx.Tx
	dbConn    *pgx.Conn
	saves     uint
	savePoint SavePoint
	opened    int
}

func Open() (*conn, error) {
	if drv == nil {
		return nil, errors.New("driver was not initialised")
	}

	drv.Lock()
	defer drv.Unlock()

	if drv.cn == nil {
		c := &conn{
			savePoint: &defaultSavePoint{},
			dbConn:    drv.dbConn,
		}

		_, err := c.beginOnce()
		if err != nil {
			return nil, err
		}

		drv.cn = c
	}

	drv.cn.opened++
	return drv.cn, nil
}

func (c *conn) Close(ctx context.Context) error {
	c.Lock()
	defer c.Unlock()

	c.opened--
	if c.opened == 0 {
		if c.tx != nil {
			c.tx.Rollback(ctx)
			c.tx = nil
		}
		drv.deleteConn()
	}

	return nil
}

func (c *conn) beginOnce() (pgx.Tx, error) {
	if c.tx == nil {
		tx, err := c.dbConn.Begin(context.Background())
		if err != nil {
			return nil, err
		}
		c.tx = tx
	}
	return c.tx, nil
}

func (c *conn) Begin(ctx context.Context) (pgx.Tx, error) {
	if c.savePoint == nil {
		return &tx{"_", c}, errors.New("savepoint is nil") // save point is not supported
	}

	c.Lock()
	defer c.Unlock()

	connTx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	c.saves++
	id := fmt.Sprintf("tx_%d", c.saves)
	_, err = connTx.Exec(ctx, c.savePoint.Create(id))
	if err != nil {
		return nil, err
	}
	return &tx{id, c}, nil
}

func (c *conn) BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error) {
	return ErrNotImplemented
}

//func (c *conn) Prepare(query string) (driver.Stmt, error) {
//	c.Lock()
//	defer c.Unlock()
//
//	tx, err := c.beginOnce()
//	if err != nil {
//		return nil, err
//	}
//
//	st, err := tx.Prepare(query)
//	if err != nil {
//		return nil, err
//	}
//	return &stmt{st: st}, nil
//}

func (c *conn) Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	return tx.Exec(ctx, sql, arguments...)
}

func (c *conn) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	return tx.Query(ctx, sql, args...)
}

func (c *conn) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginOnce()
	if err != nil {
		return nil
	}

	return tx.QueryRow(ctx, sql, args...)
}

func (c *conn) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	c.Lock()
	defer c.Unlock()

	tx, err := c.beginOnce()
	if err != nil {
		return nil, err
	}

	return tx.QueryFunc(ctx, sql, args, scans, f)
}

func (c *conn) Ping(ctx context.Context) error {
	c.Lock()
	defer c.Unlock()

	return ErrNotImplemented
}

//type stmt struct {
//	st *sql.Stmt
//}
//
//func (s *stmt) Exec(args []driver.Value) (driver.Result, error) {
//	return s.st.Exec(mapArgs(args)...)
//}
//
//func (s *stmt) NumInput() int {
//	return -1
//}
//
//func (s *stmt) Close() error {
//	return s.st.Close()
//}
//
//func (s *stmt) Query(args []driver.Value) (driver.Rows, error) {
//	rows, err := s.st.Query(mapArgs(args)...)
//	if err != nil {
//		return nil, err
//	}
//	return buildRows(rows)
//}

/*


BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error)
BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) (err error)

Prepare(ctx context.Context, name string, sql string) (sd *pgconn.StatementDescription, err error)
Deallocate(ctx context.Context, name string) error

WaitForNotification(ctx context.Context) (*pgconn.Notification, error)
IsClosed() bool


PgConn() *pgconn.PgConn
StatementCache() stmtcache.Cache
ConnInfo() *pgtype.ConnInfo
Config() *pgx.ConnConfig


SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults


*/
