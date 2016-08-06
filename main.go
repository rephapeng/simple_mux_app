package main
import (
    "server"
    "config"
)

func main() {

    /*create server */
    configuration := config.Generate()
    server := server.New(configuration.Port)
    server.StartServer()
}
