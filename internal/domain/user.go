package domain

import "time"

type User struct {
	ID        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt,omitempty" db:"updated_at,omitempty"`
}

// TableName returns the database table name for the User model
func (u *User) TableName() string {
	return "dbo.users"
}

// Columns returns the list of database columns for the User model
func (u *User) Columns() []string {
	return []string{"id", "email", "name", "created_at", "updated_at"}
}

// SelectColumns returns the list of columns to select in queries for the User model
func (u *User) SelectColumns() []string {
	return []string{"CAST(id AS NVARCHAR(36)) as id", "email", "name", "created_at", "updated_at"}
}
