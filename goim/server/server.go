package server

import (
	"fmt"
	"github.com/kerwinan/go/kit/log"
	"homework/goim/driver"
	"homework/goim/internal/workpool"
)

const (
	WORK_POOL_NUM = 10
	MAX_TASK_SIZE = 100
)

var cm *connMap

type Option func(s *Optional)

type Optional struct {
	driver []driver.Listener
}

func init() {
	err := log.Init("../config/log.yaml")
	if err != nil {
		fmt.Printf("log init err: %v", err)
	}
}

func Run(ds ...Option) (err error) {
	opt := &Optional{}
	for _, o := range ds {
		o(opt)
	}
	s := &server{
		errCh: make(chan error, 1),
	}
	cm = Init(100)
	s.workPool = workpool.NewWorkPool(WORK_POOL_NUM, MAX_TASK_SIZE)
	for _, d := range opt.driver {
		go s.accept(d)
	}
	//return s.run()
	return err
}

type server struct {
	workPool *workpool.WorkPool
	errCh    chan error
}

func (s *server) accept(ln driver.Listener) {
	for {
		conn, fd, err := ln.Accept()
		if err != nil {
			err = fmt.Errorf("accept err, %v", err)
			s.errCh <- err
			return
		}
		cm.set(fd, conn)
		// todo
		go s.run(fd)
	}
}

func (s *server) run(fd int) {
	// todo 使用连接管理池以及工作池进行处理
	conn, has := cm.get(fd)
	if !has {
		log.X.Errorf("get fd[%d] info err", fd)
		return
	}
	conn.buf = make([]byte, 1500)
	for {
		n, err := conn.netConn.Read(conn.buf[conn.offset:])
		if err != nil {
			log.X.Errorf("read err: %v")
			_ = conn.netConn.Close()
			break
		}
		conn.offset += n
		s.workPool.Add(conn.do)
	}
	return
}
