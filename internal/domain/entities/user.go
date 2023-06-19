package entities

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserSt struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Birthdate time.Time `db:"birthdate" json:"birthdate"`
}

type UserCUSt struct {
	Name      *string    `db:"name" json:"name"`
	Email     *string    `db:"email" json:"email"`
	Password  *string    `db:"password" json:"password"`
	Birthdate *time.Time `db:"birthdate" json:"birthdate"`
}

type Credentials struct {
	Password string `json:"password"`
	Login    string `json:"login"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}
