package pkg

import "net"

type RawClientMessage struct {
	Address *net.UDPAddr
	Bytes   [512]byte
}

func NewRawClientMessage(addr *net.UDPAddr, bytes [512]byte) *RawClientMessage {
	return &RawClientMessage{
		Address: addr,
		Bytes:   bytes,
	}
}
