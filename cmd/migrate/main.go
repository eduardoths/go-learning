package main

import (
	"github.com/eduardothsantos/go-learning/internal/config"
	"github.com/eduardothsantos/go-learning/migrations"
)

func main() {
	config := config.GetConfig()
	db := config.GetDBConnector()
	db.AutoMigrate(&migrations.Users{})
}
