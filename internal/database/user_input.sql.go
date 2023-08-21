// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user_input.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createUserInput = `-- name: CreateUserInput :one
INSERT INTO user_input (id, user_id, created_at, updated_at, height, weight, desired_weight, time_frame, bmi , program, curr_kcal, deficit)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
RETURNING id, user_id, created_at, updated_at, program, height, weight, desired_weight, time_frame, bmi, curr_kcal, deficit
`

type CreateUserInputParams struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Height        int32
	Weight        int32
	DesiredWeight sql.NullInt32
	TimeFrame     sql.NullInt32
	Bmi           int32
	Program       string
	CurrKcal      int32
	Deficit       sql.NullInt32
}

func (q *Queries) CreateUserInput(ctx context.Context, arg CreateUserInputParams) (UserInput, error) {
	row := q.db.QueryRowContext(ctx, createUserInput,
		arg.ID,
		arg.UserID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Height,
		arg.Weight,
		arg.DesiredWeight,
		arg.TimeFrame,
		arg.Bmi,
		arg.Program,
		arg.CurrKcal,
		arg.Deficit,
	)
	var i UserInput
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Program,
		&i.Height,
		&i.Weight,
		&i.DesiredWeight,
		&i.TimeFrame,
		&i.Bmi,
		&i.CurrKcal,
		&i.Deficit,
	)
	return i, err
}

const getUserInput = `-- name: GetUserInput :one
SELECT id, user_id, created_at, updated_at, program, height, weight, desired_weight, time_frame, bmi, curr_kcal, deficit FROM user_input where user_id = $1
`

func (q *Queries) GetUserInput(ctx context.Context, userID uuid.UUID) (UserInput, error) {
	row := q.db.QueryRowContext(ctx, getUserInput, userID)
	var i UserInput
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Program,
		&i.Height,
		&i.Weight,
		&i.DesiredWeight,
		&i.TimeFrame,
		&i.Bmi,
		&i.CurrKcal,
		&i.Deficit,
	)
	return i, err
}
