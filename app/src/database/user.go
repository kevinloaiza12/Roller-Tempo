package database

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Creation

func NewUserMap(id int, coins int, turn int) map[string]interface{} {
	return map[string]interface{}{
		"id":      id,
		"monedas": coins,
		"turno":   turn,
	}
}

func CreateNewUser(ctx context.Context, db *sql.DB, data map[string]interface{}) (bool, error) {

	id := data["id"]
	monedas := data["monedas"]
	turno := data["turno"]

	_, err := db.ExecContext(
		ctx,
		"INSERT INTO usuarios (id, monedas, turno) VALUES ($1,$2,$3)",
		id,
		monedas,
		turno,
	)

	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Utils

func usersGetQuery(ctx context.Context, db *sql.DB, attractionID int, column string) (interface{}, error) {
	var data interface{}
	query := fmt.Sprintf("SELECT %s FROM usuarios WHERE id = $1", column)
	err := db.QueryRowContext(ctx, query, attractionID).Scan(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func usersSetQuery(ctx context.Context, db *sql.DB, attractionID int, column string, value interface{}) (bool, error) {

	var query string

	switch value.(type) {
	case int:
		query = fmt.Sprintf("UPDATE usuarios SET %s = %s WHERE id = $1", column, value)
	case string:
		query = fmt.Sprintf("UPDATE usuarios SET %s = '%s' WHERE id = $1", column, value)
	}

	if err := db.QueryRowContext(ctx, query, attractionID).Err(); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// Setters

func SetUserCoinsByID(ctx context.Context, db *sql.DB, attractionID int, value string) (bool, error) {
	result, err := usersSetQuery(ctx, db, attractionID, "monedas", value)
	return result, err
}

func SetUserTurnByID(ctx context.Context, db *sql.DB, attractionID int, value string) (bool, error) {
	result, err := usersSetQuery(ctx, db, attractionID, "turno", value)
	return result, err
}

// Getters

func GetUserCoinsByID(ctx context.Context, db *sql.DB, attractionID int) (int64, error) {
	result, err := usersGetQuery(ctx, db, attractionID, "monedas")
	return result.(int64), err
}

func GetUserTurnByID(ctx context.Context, db *sql.DB, attractionID int) (int64, error) {
	result, err := usersGetQuery(ctx, db, attractionID, "turno")
	return result.(int64), err
}
