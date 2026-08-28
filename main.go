package main

import (
	"quasar/conf"
	"quasar/database"
	"quasar/server"
)

func main() {
	conf.Load()
	database.Connect()

	s := server.New()
	if err := s.Start(); err != nil {
		panic(err)
	}
}
