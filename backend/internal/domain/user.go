package domain

import "time"

type Role string

// also the table name in the DB
const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
	RoleParent  Role = "parent"
	RoleAdmin   Role = "admin"
)

// User is the core entity — db tags for sqlx, json tags for direct marshaling
type User struct {
	// ID is unique
	ID string `db:"id"         json:"id"`
	// Username is unique
	Username string `db:"username"       json:"username"`
	// Email is unique
	Email        string    `db:"email"      json:"email"`
	HashPassword string    `db:"hash_password"   json:"-"`
	Role         Role      `db:"role"       json:"role"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type Parent struct {
	User
}

type Student struct {
	User
}

type Teacher struct {
	User
}

type Admin struct {
	User
}
