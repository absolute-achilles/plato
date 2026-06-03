package domain

import "time"

type Role string

// also the table name in the DB
const (
	RoleStudent  Role = "student"
	RoleTeacher  Role = "teacher"
	RoleGuardian Role = "guardian"
	RoleAdmin    Role = "admin"
)

// User is the core entity — db tags for sqlx, json tags for direct marshaling
type User struct {
	ID        string    `db:"id"         json:"id"`
	Name      string    `db:"name"       json:"name"`
	Email     string    `db:"email"      json:"email"`
	Password  string    `db:"password"   json:"-"`
	Role      Role      `db:"role"       json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at,omitempty"`
}

type Guardian struct {
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
