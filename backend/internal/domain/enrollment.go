package domain

import "time"

type Enrollment struct {
	ID         string    `db:"id"`
	StudentID  string    `db:"student_id"`
	CourseID   string    `db:"course_id"`
	EnrolledAt time.Time `db:"enrolled_at"`
}
