package domain

type Role string

const (
	RoleStudent Role = "student"
	RoleTeacher Role = "teacher"
	RoleParent  Role = "parent"
	RoleAdmin   Role = "admin"
)

type User struct {
	// ID is unique
	ID string `db:"id"         json:"id"`
	// Username is unique
	Username string `db:"username"       json:"username"`
	// Email is unique
	Email        string `db:"email"      json:"email"`
	HashPassword string `db:"hash_password"   json:"-"`
	Role         Role   `db:"role"       json:"role"`
	// CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type ParentRelationship string

const (
	ParentRelationshipFather   ParentRelationship = "father"
	ParentRelationshipMother   ParentRelationship = "mother"
	ParentRelationshipGuardian ParentRelationship = "guardian"
)

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
