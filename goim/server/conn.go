package server

import (
	"github.com/kerwinan/go/kit/log"
	"homework/goim/internal/packet"
	"net"
	"time"
)

const (
	BUFF_MAX_SIZE = 1500
)

type connection struct {
	fd      int
	netConn net.Conn
	buf     []byte
	offset  int
}

func (c *connection) do() {
	pkt, remainData, err := packet.Decode(c.buf[:c.offset])
	if err != nil {
		log.X.Errorf("decode err: %v", err)
		_ = c.netConn.Close()
		return
	}
	if len(remainData) != 0 {
		copy(c.buf, remainData)
		c.offset = len(remainData)
	} else {
		c.offset = 0
	}
	log.X.Debugf("recv data: %+v", *pkt)
	// todo somethings

	str := "hello"
	pktResp := packet.IMPacket{}
	pktResp.HeaderLen = int16(packet.HeaderLength)
	pktResp.Sequence = pkt.Sequence + 1
	pktResp.Version = pkt.Version
	pktResp.Body = append(pktResp.Body, []byte(str)...)

	enData, err := packet.Encode(pktResp)
	if err != nil {
		log.X.Errorf("packer encode err, %+v", err)
		return
	}
	_ = c.netConn.SetWriteDeadline(time.Now().Add(time.Second * 3))
	n, err := c.netConn.Write(enData)
	if err != nil {
		log.X.Errorf("send data err, %+v", err)
		return
	}
	log.X.Debugf("send %d bytes data success", n)
}
