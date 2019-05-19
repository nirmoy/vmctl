package dummy

import (
	"errors"
	"math/rand"

	"github.com/google/uuid"
)

type Server struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ServerID struct {
	ID string `json:"id"`
}

type ServerStatus struct {
	CpuUtilization int `json:"cpuUtilization"`
}

var blackListedNames = []string{"blacklisted", "blacklisted1", "blacklisted2"}
var prohibitedNames = []string{"prohibited", "prohibited1", "prohibited2"}
var allServer = []Server{}

func isBlacklistedVMName(name string) bool {
	for _, blackListedName := range blackListedNames {
		if blackListedName == name {
			return true
		}
	}
	return false
}

func isProhibitedVMName(name string) bool {
	for _, prohibitedName := range prohibitedNames {
		if prohibitedName == name {
			return true
		}
	}
	return false
}

func CreateServer(name string) (Server, error) {
	if isBlacklistedVMName(name) {
		return Server{}, errors.New("blacklisted VM name")
	}

	id := uuid.New()
	server := Server{
		Name: name,
		ID:   id.String(),
	}
	allServer = append(allServer, server)
	return server, nil
}

func DeleteServerByUUID(uuid string) (bool, error) {
	if !IsExistServerByUUID(uuid) {
		return false, errors.New("Not Found")
	}

	index := -1
	for i, server := range allServer {
		if server.ID == uuid {
			index = i
			break
		}
	}

	// index can't be -1 as we already check the server
	allServer = append(allServer[:index], allServer[index+1:]...)

	return true, nil
}

func GetAllServer() ([]Server, error) {
	return allServer, nil
}

func GetServerByUUID(uuid string) (Server, error) {
	for _, server := range allServer {
		if server.ID == uuid {
			return server, nil
		}
	}
	// no error for dummy
	return Server{}, nil
}

func GetServerStatusByUUID(uuid string) (ServerStatus, error) {
	for _, server := range allServer {
		if server.ID == uuid {
			status := ServerStatus{
				CpuUtilization: rand.Intn(100),
			}
			return status, nil
		}
	}
	// no error for dummy
	return ServerStatus{}, nil
}

func GetServerByName(name string) Server {
	for _, server := range allServer {
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

func IsProhibitedServer(name string) bool {
	if IsExistServerByName(name) {
		if isProhibitedVMName(name) {
			return true
		}
	}

	return false
}

func IsExistServerByUUID(uuid string) bool {
	server, err := GetServerByUUID(uuid)

	if err != nil {
		return false
	}

	if (len(server.ID) + len(server.Name)) == 0 {
		return false
	}

	return true
}
