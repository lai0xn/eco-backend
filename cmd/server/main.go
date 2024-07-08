package main

import (
	"github.com/lai0xn/squid-tech/config"
	_ "github.com/lai0xn/squid-tech/docs"
	"github.com/lai0xn/squid-tech/internal/server"
	"github.com/lai0xn/squid-tech/prisma"
)

// @title			Squid Tech API
// @version		1.0
// @description	backend of the event management app.
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	config.Load()
	s := server.NewServer(":8080")
	prisma.Connect()
	s.Run()
}
