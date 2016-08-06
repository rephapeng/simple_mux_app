package server

import (
    "fmt"
    "log"
    "net/http"
    "router"
)

type Server struct {
    Port string
}

func New(port string) (server *Server) {
    return &Server{
        Port: port,
    }
}

/*instance server*/
func (server *Server) StartServer() {
    router := router.NewRouter()

    log.Fatal(http.ListenAndServe(":"+server.Port, router))
}

func (server *Server) Print() {
    fmt.Println("Server port:", server.Port)
}
