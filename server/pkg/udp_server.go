package pkg

import (
	"log"
	"net"
	"os"
)

type UDPServer struct {
	HostPort       string
	Address        *net.UDPAddr
	Connection     *net.UDPConn
	Channel        chan<- *RawClientMessage
	LoggingEnabled bool
	Logger         *log.Logger
}

func NewUDPServer(hostPort string) (*UDPServer, error) {
	addr, err := net.ResolveUDPAddr("udp", hostPort)
	if err != nil {
		return nil, err
	}

	return &UDPServer{
		HostPort: hostPort,
		Address:  addr,
		Channel:  make(chan<- *RawClientMessage),
		Logger:   log.New(os.Stdout, "UDPSERVER: ", log.LstdFlags),
	}, nil
}

func (u *UDPServer) Listen() error {
	conn, err := net.ListenUDP("udp", u.Address)
	if err != nil {
		return err
	}

	u.Connection = conn

	for {
		var buf [512]byte
		_, addr, readErr := conn.ReadFromUDP(buf[0:])
		if readErr != nil {
			return readErr
		}

		if u.LoggingEnabled {
			u.Logger.Printf("[FROM CLIENT]: Addr(%s) | Size (%d)", addr.String(), len(buf))
		}

		u.Channel <- NewRawClientMessage(addr, buf)
	}
}

func (u *UDPServer) SetLogging(trueOrFalse bool) {
	u.LoggingEnabled = trueOrFalse
}
