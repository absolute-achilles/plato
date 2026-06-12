package domain

import "time"

type FileType string

const (
	FileTypeImagePng  FileType = "image/png"
	FileTypeImageJpg  FileType = "image/jpg"
	FileTypeImageGif  FileType = "image/gif"
	FileTypeImageWebp FileType = "image/webp"

	FileTypeVideoMp4 FileType = "video/mp4"
)

type ModuleContentType string

const (
	ModuleContentTypeLesson     string = "lesson"
	ModuleContentTypeAssignment string = "assignment"
)

type Module struct {
	ID           string    `db:"id"`
	CourseID     string    `db:"course_id"`
	Name         string    `db:"name"`
	Position     float64   `db:"position"`
	IsPusblished bool      `db:"is_published"`
	UnlockDate   time.Time `db:"unlock_date"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type ModuleContent struct {
	ID                string            `db:"id"`
	ModuleID          string            `db:"module_id"`
	Title             string            `db:"title"`
	ModuleContentType ModuleContentType `db:"position"`
	BodyContent       string            `db:"title"`
	Position          float64           `db:"position"`
	IsPusblished      bool              `db:"is_published"`
	CreatedAt         time.Time         `db:"created_at"`
	UpdatedAt         time.Time         `db:"updated_at"`
}

type ContentAttachment struct {
	ID              string   `db:"id"`
	ModuleContentID string   `db:"module_content_id"`
	Name            string   `db:"name"`
	URL             string   `db:"url"`
	SizeBytes       int64    `db:"size_bytes"`
	FileType        FileType `db:"type"`
}
