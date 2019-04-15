package data

type Cred struct {
	Name string `json:"name" db:"name"`
	Pass string `json:"pass"  db:"pass"`
}

type User struct {
	Cred
	Email string `json:"email"  db:"email"`
	Role  string `json:"user_role"  db:"user_role"`
}
