package data

type Cred struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

type User struct {
	Cred
	Email string `json:"name"`
	Role  string `json:"role"`
}
