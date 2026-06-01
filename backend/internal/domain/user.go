package domain

import "time"

type Role string

const (
	RoleStudent  Role = "student"
	RoleTeacher  Role = "teacher"
	RoleGuardian Role = "guardian"
	RoleAdmin    Role = "admin"
)

// User is the core entity — db tags for sqlx, json tags for direct marshaling
type User struct {
	ID    int64  `db:"id"         json:"id"`
	Name  string `db:"name"       json:"name"`
	Email string `db:"email"      json:"email"`
	Role  Role   `db:"role"       json:"role"`
	// should omit this i think
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
