package server

import (
	"net"
	"sync"
	"time"
)

type connMap struct {
	mutex sync.RWMutex
	conns []connection
	m     map[int]int64
}

func Init(size int64) *connMap {
	m := &connMap{
		m:     make(map[int]int64, size),
		conns: make([]connection, size),
	}
	for i := int64(0); i < size; i++ {
		m.conns[i] = connection{
			buf: make([]byte, BUFF_MAX_SIZE),
		}
	}
	return m
}

func (m *connMap) set(fd int, conn net.Conn) {
	// todo
	m.mutex.RLock()
	c := connection{
		fd:      fd,
		netConn: conn,
	}
	// todo 优化
	t := time.Now().Nanosecond()
	m.conns[t] = c
	m.mutex.RUnlock()
}

func (m *connMap) get(fd int) (*connection, bool) {
	m.mutex.RLock()
	index, has := m.m[fd]
	if !has {
		m.mutex.RUnlock()
		return nil, false
	}
	c := &m.conns[index]
	m.mutex.RUnlock()
	return c, true
}
