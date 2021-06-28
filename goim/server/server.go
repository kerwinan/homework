package server

import (
	"fmt"
	"homework/goim/driver"
	"net"
)

type Option func(s *Optional)

type Optional struct {
	driver []driver.Listener
}

func Run(ds ...Option) (err error) {
	opt := &Optional{}
	for _, o := range ds {
		o(opt)
	}
	s := &server{}
	for _, d := range opt.driver {
		go s.accept(d)
	}
	//return s.run()
	return err
}

type server struct {
	errCh chan error
}

func (s *server) accept(ln driver.Listener) {
	for {
		conn, fd, err := ln.Accept()
		if err != nil {
			err = fmt.Errorf("accept err, %v", err)
			s.errCh <- err
			return
		}
		// todo
		go run(conn, fd)
	}
}

func (s *server) run() (err error) {
	// todo 使用连接管理池以及工作池进行处理
	return nil
}

func run(conn net.Conn, fd int) {
	for {

	}
	return
}
