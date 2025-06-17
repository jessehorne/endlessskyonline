package pkg

import (
	"fmt"
	"log"
	"os"
)

type GameServer struct {
	UDPServer      *UDPServer
	LoggingEnabled bool
	Logger         *log.Logger
}

func NewGameServer(hostPort string) (*GameServer, error) {
	// initialize udp server
	udpServer, err := NewUDPServer(hostPort)
	if err != nil {
		return nil, err
	}
	udpServer.SetLogging(true)
	return &GameServer{
		UDPServer: udpServer,
		Logger:    log.New(os.Stdout, "GAMESERVER: ", log.LstdFlags),
	}, nil
}

func (gs *GameServer) Start() error {
	gs.Log(fmt.Sprintf("Starting GameServer | Addr (%s)", gs.UDPServer.Address.String()))

	// listen for udp packets
	err := gs.UDPServer.Listen()
	if err != nil {
		return err
	}
	return nil
}

func (gs *GameServer) SetLogging(trueOrFalse bool) {
	gs.LoggingEnabled = trueOrFalse
	gs.UDPServer.SetLogging(trueOrFalse)
}

func (gs *GameServer) Log(l string) {
	if gs.LoggingEnabled {
		gs.Logger.Println(l)
	}
}
