package db

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const defaultDialect = "postgres"

type User struct {
	ID             string
	Name           string
	DocumentNumber int64
	Email          string
}

type DB struct {
	*sql.DB
}

func NewDB(connectionS string) *DB {
	databaseConn, err := sql.Open(defaultDialect, connectionS)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	return &DB{DB: databaseConn}
}

func (db *DB) InsertUser(user *User) error {
	insertStmt := `
        INSERT INTO users (id, name, document_number, email)
        VALUES ($1, $2, $3, $4)
    `
	_, err := db.Exec(insertStmt, uuid.New().String(), user.Name, user.DocumentNumber, user.Email)
	return err
}

func (db *DB) FindUserByEmail(value string) (*User, error) {
	var user User
	query := `
        SELECT id, name, document_number, email
        FROM users
        WHERE lower(name) = $1 OR email = $2
    `

	row := db.QueryRow(query, strings.ToLower(value), value)
	err := row.Scan(&user.ID, &user.Name, &user.DocumentNumber, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (db *DB) FindAll() (users []User, err error) {
	query := `
        SELECT id, name, document_number, email
        FROM users
        WHERE 1=1 limit 10
    `

	rows, err := db.Query(query)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Name, &user.DocumentNumber, &user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
