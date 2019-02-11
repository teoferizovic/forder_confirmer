package model

type User struct {
	ID int64				`json:"id,omitempty"`
	Email string			`json:"email,omitempty"`
	Password string			`json:"password,omitempty"`
	Role_id int64			`json:"role_id,omitempty"`
	Created_at string		`json:"created_at,omitempty"`
	Updated_at string		`json:"updated_at,omitempty"`
}
