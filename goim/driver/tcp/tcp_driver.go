package tcp

import (
	"homework/goim/driver"
	"net"
	"reflect"
	"time"
)

//NewListener new listener
func NewListener(addr string) (ln driver.Listener, err error) {
	tcpLn, err := net.Listen("tcp", addr)
	if err != nil {
		return
	}
	ln = &listener{
		ln: tcpLn,
	}
	return
}

type listener struct {
	ln net.Listener
}

func (l *listener) Accept() (conn net.Conn, fd int, err error) {
	conn, err = l.ln.Accept()
	if err != nil {
		return
	}
	fd = int(reflect.Indirect(reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").FieldByName("fd")).FieldByName("pfd").FieldByName("Sysfd").Int())
	return
}

func (l *listener) Close() error {
	return l.ln.Close()
}

func (l *listener) Addr() net.Addr {
	return l.ln.Addr()
}

//NewDialer new dialer
func NewDialer(addr string) driver.Dialer {
	return &dialer{
		addr: addr,
	}
}

type dialer struct {
	addr string
}

func (d *dialer) Dial(timeout time.Duration) (conn net.Conn, fd int, err error) {
	conn, err = net.DialTimeout("tcp", d.addr, timeout)
	if err != nil {
		return
	}
	fd = int(reflect.Indirect(reflect.Indirect(reflect.ValueOf(conn)).FieldByName("conn").FieldByName("fd")).FieldByName("pfd").FieldByName("Sysfd").Int())
	return
}
