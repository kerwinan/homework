package driver

import (
	"net"
	"time"
)

//Listener listener
type Listener interface {
	Accept() (conn net.Conn, fd int, err error)
	Close() error
	Addr() net.Addr
}

//Dialer dialer
type Dialer interface {
	Dial(timeout time.Duration) (conn net.Conn, fd int, err error)
}