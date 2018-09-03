package server

import (
	"fmt"
	"log"
	"net/http"
)

// ServerConfig is used for Server Configurations
type ServerConfig struct {
	HOST string
	PORT string
}

// NewServer is used to set ServerConfig
func NewServer(host, port string) ServerConfig {
	return ServerConfig{HOST: host, PORT: port}
}

// Start will start the server
func (config ServerConfig) Start() {
	// setAPIs()
	fmt.Println("Server Listening on : ", config.PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.HOST, config.PORT), nil))
}
