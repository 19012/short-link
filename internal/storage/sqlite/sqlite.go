package sqlite

import (
	"19012/short-link/internal/storage"
	"database/sql"
	"errors"
	"fmt"

	"github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.sglite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	stmt, err := db.Prepare(`
    create table if not exists url(
    id integer primary key,
    alias text not null unique,
    url text not null);
    create index if not exists idx_alias on url(alias);
    `)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec()
	if err != nil {
		return nil, fmt.Errorf("%s:%w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveURL(urlToSave string, alias string) (int64, error) {
	const op = "storage.sqlite.SaveURL"

	stmt, err := s.db.Prepare("insert into url(url, alias) values(?,?)")
	if err != nil {
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	res, err := stmt.Exec(urlToSave, alias)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return 0, fmt.Errorf("%s:%w", op, storage.ErrURLExists)
		}
		return 0, fmt.Errorf("%s:%w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: failed to get last insert id: %w", op, err)
	}

	return id, nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const op = "storage.sqlite.GetURL"

	stmt, err := s.db.Prepare("select url from url where alias = ?")
	if err != nil {
		return "", fmt.Errorf("%s:%w", op, err)
	}

	row := stmt.QueryRow(alias)
	var res string
	if err := row.Scan(&res); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("%s:%w", op, storage.ErrURLNotFound)
		}
		return "", fmt.Errorf("%s:%w", op, err)
	}
	return res, nil
}

func (s *Storage) DeleteURL(alias string) (bool, error) {
	const op = "storage.sqlite.DeleteURL"

	stmt, err := s.db.Prepare("delete from url where alias = ?")
	if err != nil {
		return false, fmt.Errorf("%s:%w", op, err)
	}

	res, err := stmt.Exec()
	if err != nil {
		return false, fmt.Errorf("%s:%w", op, err)
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("%s:%w", op, err)
	}

	if affected > 0 {
		return true, nil
	}

	return false, nil
}
