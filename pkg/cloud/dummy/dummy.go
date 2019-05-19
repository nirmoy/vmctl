package dummy

import (
	"github.com/google/uuid"
)

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var AllServers = []Server{}

func CreateServer(name string) (Server, error) {
	id := uuid.New()
	server := Server{
		Name: name,
		ID:   id.String(),
	}
	AllServers = append(AllServers, server)
	return server, nil
}

func GetAllServers() []Server {
	return AllServers
}
