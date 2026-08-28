package main

import (
	"quasar/conf"
	"quasar/server"
)

func main() {
	conf.Load()

	s := server.New()
	if err := s.Start(); err != nil {
		panic(err)
	}
}
