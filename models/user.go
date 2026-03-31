package models

type User struct {
	UserID   string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

type Token struct{
	TokenId int 
	Token string
	UserID string
	ExpiresAT int
}