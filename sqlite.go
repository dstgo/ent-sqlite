package ent_sqlite

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/glebarez/go-sqlite"
)

func init() {
	sql.Register("sqlite3", Driver{&sqlite.Driver{}})
}

// Driver is sqlite driver wrapper
type Driver struct {
	*sqlite.Driver
}

func (d Driver) Open(dsn string) (driver.Conn, error) {
	conn, err := d.Driver.Open(dsn)
	if err != nil {
		return nil, err
	}

	// ent sqlite must be enabled with fk
	if execer, ok := conn.(driver.ExecerContext); ok {
		_, err := execer.ExecContext(context.Background(), "PRAGMA foreign_keys = on;", nil)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unsupported driver connection")
	}

	return conn, nil
}
