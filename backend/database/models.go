// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Container struct {
	ID      int64
	Name    pgtype.Text
	LocalID pgtype.Text
	Image   pgtype.Text
	Status  pgtype.Text
}

type Flow struct {
	ID          int64
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	Name        pgtype.Text
	Status      pgtype.Text
	ContainerID pgtype.Int8
}

type Task struct {
	ID        int64
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
	Type      pgtype.Text
	Status    pgtype.Text
	Args      []byte
	Results   pgtype.Text
	FlowID    pgtype.Int8
	Message   pgtype.Text
}
