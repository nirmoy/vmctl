package dummy

import (
	"github.com/google/uuid"
)

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ServerID struct {
	ID string `json:"id"`
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

func GetServerByID(uuid string) Server {
	for _, server := range AllServers {
		if server.ID == uuid {
			return server
		}
	}
	return Server{}
}

func GetServerByName(name string) Server {
	for _, server := range AllServers {
		if server.Name == name {
			return server
		}
	}
	return Server{}
}

func IsExistServerByName(name string) bool {
	server := GetServerByName(name)
	if (len(server.ID) + len(server.Name)) == 0 {
		return false
	}

	return true
}

func IsExistServerByID(uuid string) bool {
	server := GetServerByID(uuid)
	if (len(server.ID) + len(server.Name)) == 0 {
		return false
	}

	return true
}
