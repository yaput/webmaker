package database

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

func NewDBConnection(conf Config) (con *Connection, err error) {
	if conf.Host == "" || conf.Port == "" || conf.User == "" ||
		conf.Password == "" || conf.Database == "" {
		err = errors.New("All fields must be set")
		return
	}

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		conf.User, conf.Password, conf.Database, conf.Host, conf.Port))
	if err != nil {
		err = fmt.Errorf("Couldn't open connection to postgre database got: %+v", err)
		return
	}

	if err = db.Ping(); err != nil {
		err = fmt.Errorf("Couldn't ping postgre database got: %+v", err)
		return
	}

	return &Connection{
		cfg: conf,
		Db:  db,
	}, nil
}

func (c *Connection) GetHost() string {
	return c.cfg.Host
}

func (c *Connection) GetDatabase() string {
	return c.cfg.Database
}

func (c *Connection) GetUser() string {
	return c.cfg.User
}

func (c *Connection) Close(err error) {
	if c.Db == nil {
		return
	}

	if err = c.Db.Close(); err != nil {
		err = fmt.Errorf("Errored closing database connection got: %+v", err)
	}

	return
}

func (c *Connection) SelectRow(query SelectQuery) *SelectQuery {
	fields := strings.Join(query.Field, ",")
	query.conn = c.Db
	query.queryExec = fmt.Sprintf(`SELECT %s FROM %s`, fields, query.Table)
	return &query
}

func (q *SelectQuery) Where(condition string) *SelectQuery {
	var cond string
	if !strings.Contains(q.queryExec, "WHERE") {
		cond += "WHERE"
	} else {
		cond += "AND"
	}
	cond += condition
	q.queryExec += cond
	return q
}

func (q *SelectQuery) OrWhere(condition string) *SelectQuery {
	var cond string
	if !strings.Contains(q.queryExec, "WHERE") {
		cond += "WHERE"
	} else {
		cond += "OR"
	}
	cond += condition
	q.queryExec += cond
	return q
}

func (q *SelectQuery) Get() (r *sql.Rows, err error) {
	if len(strings.Split(q.queryExec, ";")) > 1 {
		err = errors.New("Multiple Query! Not Allowed")
		return
	}
	r, err = q.conn.Query(q.queryExec)
	if err != nil {
		err = fmt.Errorf("Exec query got: %+v", err)
		return
	}
	return
}
