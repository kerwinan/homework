package packet

import (
	"fmt"
	"homework/goim/internal"
)

const (
	OpHandshake       = int32(0) // handshake
	OpHandshakeReply  = int32(1) // handshake reply
	OpHeartbeat       = int32(2) // heartbeat
	OpHeartbeatReply  = int32(3) // heartbeat reply
	OpSendMsg         = int32(4) // send message
	OpSendMsgReply    = int32(5) // send message reply
	OpDisconnectReply = int32(6) // connection disconnect reply
	OpAuth            = int32(7) // auth connnect
	OpAuthReply       = int32(8) // auth connect reply
	OpRawBatch        = int32(9) // batch message for websocket
)

var (
	_headerLength = 4 + 2 + 2 + 4 + 4
)

type IMPacket struct {
	PacketLen int32  // header + body length
	HeaderLen int16  // protocol header length
	Version   int16  // protocol version
	Operation int32  // operation for request
	Sequence  int32  // sequence number chosen by client
	Body      []byte // binary body bytes
}

func Encode(pkt IMPacket) ([]byte, error) {
	data := make([]byte, pkt.PacketLen)
	var err error

	bytes := internal.Int32ToBytes(pkt.PacketLen)
	index := uint64(0)
	copy(data[index:], bytes)
	index += uint64(len(bytes))

	bytes = internal.Int16ToBytes(pkt.HeaderLen)
	copy(data[index:], bytes)
	index += uint64(len(bytes))

	bytes = internal.Int16ToBytes(pkt.Version)
	copy(data[index:], bytes)
	index += uint64(len(bytes))

	bytes = internal.Int32ToBytes(pkt.Operation)
	copy(data[index:], bytes)
	index += uint64(len(bytes))

	bytes = internal.Int32ToBytes(pkt.Sequence)
	copy(data[index:], bytes)
	index += uint64(len(bytes))

	copy(data[index:], pkt.Body)
	return data, err
}

func Decode(data []byte) (pkt *IMPacket, remainData []byte, err error) {
	remainData = data
	pkt = &IMPacket{}
	pkt.PacketLen = internal.BytesToInt32(data[0:4])
	if pkt.PacketLen > int32(len(data)) {
		err = fmt.Errorf("data not enough")
		return
	}
	pkt.HeaderLen = internal.BytesToInt16(data[4:6])
	pkt.Version = internal.BytesToInt16(data[6:8])
	pkt.Operation = internal.BytesToInt32(data[8:12])
	pkt.Sequence = internal.BytesToInt32(data[12:16])
	bodyLen := pkt.PacketLen - int32(pkt.HeaderLen)
	pkt.Body = make([]byte, bodyLen)
	copy(pkt.Body, data[_headerLength:int32(_headerLength)+bodyLen])

	remainData = make([]byte, int32(len(data))-pkt.PacketLen)
	copy(remainData, data[pkt.PacketLen:])
	return
}
