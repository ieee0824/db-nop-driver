package dbnop

import (
	"database/sql"
	"database/sql/driver"
)

func init() {
	sql.Register("dbnop", &Driver{})
}

type Result struct{}

func (r *Result) LastInsertId() (int64, error) {
	return 0, nil
}

func (r *Result) RowsAffected() (int64, error) {
	return 0, nil
}

type Rows struct {
}

func (r *Rows) Columns() []string {
	return []string{}
}

func (r *Rows) Close() error {
	return nil
}

func (r *Rows) Next(dest []driver.Value) error {
	return sql.ErrNoRows
}

type Stmt struct{}

func (s *Stmt) Close() error {
	return nil
}

func (s *Stmt) NumInput() int {
	return 0
}

func (s *Stmt) Exec(args []driver.Value) (driver.Result, error) {
	return &Result{}, nil
}

func (s *Stmt) Query(args []driver.Value) (driver.Rows, error) {
	return &Rows{}, nil
}

type Tx struct {
}

func (t *Tx) Commit() error {
	return nil
}

func (t *Tx) Rollback() error {
	return nil
}

type Conn struct {
}

func (c *Conn) Prepare(query string) (driver.Stmt, error) {
	return &Stmt{}, nil
}

func (c *Conn) Close() error {
	return nil
}

func (c *Conn) Begin() (driver.Tx, error) {
	return &Tx{}, nil
}

type Driver struct {
}

func (d *Driver) Open(name string) (driver.Conn, error) {
	return &Conn{}, nil
}
