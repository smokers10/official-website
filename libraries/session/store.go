package session

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mysql"
	config "github.com/smokers10/official-website.git/libraries"
)

type sessions struct{}

func Store() *sessions {
	return &sessions{}
}

func (s *sessions) MYSQL() *session.Store {
	config := config.ReadConfig().Database
	mysqlStorage := mysql.New(mysql.Config{
		Host:       config.Host,
		Port:       config.Port,
		Username:   config.Username,
		Password:   config.Password,
		Database:   config.Database,
		Table:      config.SessionTable,
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	CDConfig := session.ConfigDefault

	CDConfig.Storage = mysqlStorage

	return session.New(CDConfig)
}
