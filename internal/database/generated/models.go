// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package database

import (
	"database/sql"
)

type Tile struct {
	ID    string
	Title string
	Type  string
	Link  sql.NullString
	Image sql.NullString
}
