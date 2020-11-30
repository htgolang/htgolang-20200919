package socket

import (
	"net"
)

// Server for remote user manager
func Server() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		panic(err)
	}
	conn, errA := listener.Accept()
	if errA != nil {
		panic(errA)
	}

}
