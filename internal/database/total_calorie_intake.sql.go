// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: total_calorie_intake.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTotalCalorieIntake = `-- name: CreateTotalCalorieIntake :one
INSERT INTO total_calorie_intake(id, user_id, calories, program, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, created_at, updated_at, calories, user_id, program
`

type CreateTotalCalorieIntakeParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Calories  int32
	Program   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateTotalCalorieIntake(ctx context.Context, arg CreateTotalCalorieIntakeParams) (TotalCalorieIntake, error) {
	row := q.db.QueryRowContext(ctx, createTotalCalorieIntake,
		arg.ID,
		arg.UserID,
		arg.Calories,
		arg.Program,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i TotalCalorieIntake
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Calories,
		&i.UserID,
		&i.Program,
	)
	return i, err
}

const deleteRedundantRows = `-- name: DeleteRedundantRows :exec
DELETE FROM total_calorie_intake WHERE calories = 0
`

func (q *Queries) DeleteRedundantRows(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteRedundantRows)
	return err
}

const getMostRecentUserKcal = `-- name: GetMostRecentUserKcal :one
SELECT id, created_at, updated_at, calories, user_id, program FROM total_calorie_intake WHERE user_id = $1 ORDER BY created_at DESC LIMIT 1
`

func (q *Queries) GetMostRecentUserKcal(ctx context.Context, userID uuid.UUID) (TotalCalorieIntake, error) {
	row := q.db.QueryRowContext(ctx, getMostRecentUserKcal, userID)
	var i TotalCalorieIntake
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Calories,
		&i.UserID,
		&i.Program,
	)
	return i, err
}

const getTotalCalories = `-- name: GetTotalCalories :many
SELECT id, created_at, updated_at, calories, user_id, program FROM total_calorie_intake WHERE user_id = $1
`

func (q *Queries) GetTotalCalories(ctx context.Context, userID uuid.UUID) ([]TotalCalorieIntake, error) {
	rows, err := q.db.QueryContext(ctx, getTotalCalories, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TotalCalorieIntake
	for rows.Next() {
		var i TotalCalorieIntake
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Calories,
			&i.UserID,
			&i.Program,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
