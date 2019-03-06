package model

type Payment struct {
	ID int64			`json:"id,gorm:"primary_key""`
	Name string			`json:"name,omitempty"`
}
