package server

import (
	"quasar/conf"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

var Config = &conf.C

type Server struct {
	Host      string
	Port      int
	appServer *fiber.App
}

// NewCustom creates and returns a new Server instance with the specified host, port, and Fiber app configuration.
func NewCustom(host string, port int, appConfig fiber.Config) *Server {
	return &Server{
		Host:      host,
		Port:      port,
		appServer: fiber.New(appConfig),
	}
}

// New creates and returns a new Server instance configured from the config file.
func New() *Server {
	host := Config.Server.Generic.Host
	port := Config.Server.Generic.Port
	appConfig := fiber.Config{
		AppName:       Config.Server.Instance.Name,
		CaseSensitive: Config.Server.Instance.CaseSensitive,
		StrictRouting: Config.Server.Instance.StrictRouting,
	}

	return NewCustom(host, port, appConfig)
}

// StartCustom initializes and begins listening for incoming requests on the server using the specified listen configuration.
func (server *Server) StartCustom(listenConfig fiber.ListenConfig) error {
	address := server.Host + ":" + strconv.Itoa(server.Port)
	return server.appServer.Listen(address, listenConfig)
}

// Start initializes and begins listening for incoming requests on the server using the listen configuration from the config file.
func (server *Server) Start() error {
	listenConfig := fiber.ListenConfig{
		EnablePrefork:     Config.Server.Listen.EnablePrefork,
		EnablePrintRoutes: Config.Server.Listen.EnablePrintRoutes,
	}
	return server.StartCustom(listenConfig)
}
