// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Bar struct {
	ID   int32
	Name pgtype.Text
}