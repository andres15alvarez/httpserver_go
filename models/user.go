package models

type User struct {
	tableName struct{} `pg:"server_user"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
}
