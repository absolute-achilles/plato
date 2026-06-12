package domain

import "time"

type Module struct {
	ID           string    `db:"id"`
	CourseID     string    `db:"course_id"`
	Name         string    `db:"name"`
	Position     float64   `db:"position"`
	IsPusblished bool      `db:"is_published"`
	UnlockDate   time.Time `db:"unlock_date"`
	UpdatedAt    time.Time `db:"updated_at"`
}
