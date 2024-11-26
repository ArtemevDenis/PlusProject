package climes

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Climes struct {
	jwt.StandardClaims
	UID    int64
	Email  string
	Exp    time.Time
	App_id int64
}
