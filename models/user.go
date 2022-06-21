package models

type UserBase struct {
	tableName struct{} `pg:"server_user"`
	Id        int64    `json:"id" pg:",pk"`
	Name      string   `json:"name"`
	Email     string   `json:"email"`
}

type User struct {
	tableName struct{} `pg:"server_user"`
	UserBase
	Password string `json:"password"`
}

type UserCreated struct {
	tableName struct{} `pg:"server_user"`
	UserBase
}
