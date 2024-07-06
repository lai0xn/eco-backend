package main

import (
	_ "github.com/lai0xn/squid-tech/docs"
	"github.com/lai0xn/squid-tech/internal/server"
)

// @title			Squid Tech API
// @version		1.0
// @description	backend of the event management app.
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	s := server.NewServer(":8080")
	s.Run()
}
