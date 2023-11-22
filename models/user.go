// models/user.go
package models

type User struct {
	ID    int    `json:"id" xorm:"'id' pk autoincr"`
	Name  string `json:"name" xorm:"'name'"`
	Email string `json:"email" xorm:"'email'"`
}
