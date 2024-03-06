// Todo Model
package models

type Todo struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}
