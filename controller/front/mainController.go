package front

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/middleware/session"
	db "github.com/smokers10/official-website.git/libraries/database"
	bcrypt "github.com/smokers10/official-website.git/libraries/encryption"
	jwt "github.com/smokers10/official-website.git/libraries/jsonwebtoken"
	sessionStore "github.com/smokers10/official-website.git/libraries/session"
)

type MainController struct{}

func (bf *MainController) database() *sql.DB {
	mysql, err := db.Databases().MYSQLConnection()
	if err != nil {
		panic(err)
	}
	return mysql
}

func (mc *MainController) bcrypt() *bcrypt.BcryptEncryption {
	return bcrypt.Bcrypt()
}

func (mc *MainController) session() *session.Store {
	return sessionStore.Store().MYSQL()
}

func (mc *MainController) jwt(payload *jwt.Payload) *jwt.JwonWebToken {
	return jwt.JsonWebToken(payload)
}
