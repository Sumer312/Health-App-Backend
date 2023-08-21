// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DailyCalorieIntake struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Calories  int32
	UserID    uuid.UUID
}

type TotalCalorieIntake struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Calories  int32
	UserID    uuid.UUID
	Program   string
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     sql.NullString
	Password  string
}

type UserInput struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Program       string
	Height        int32
	Weight        int32
	DesiredWeight sql.NullInt32
	TimeFrame     sql.NullInt32
	Bmi           int32
	CurrKcal      int32
	Deficit       sql.NullInt32
}
