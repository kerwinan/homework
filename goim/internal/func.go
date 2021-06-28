package internal

import (
	"bytes"
	"encoding/binary"
)

func Int32ToBytes(num int32) []byte {
	var buffer bytes.Buffer
	_ = binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}

func Int16ToBytes(num int16) []byte {
	var buffer bytes.Buffer
	_ = binary.Write(&buffer, binary.BigEndian, num)
	return buffer.Bytes()
}

func BytesToInt32(buf []byte) int32 {
	var num int32
	buffer := bytes.NewBuffer(buf)
	_ = binary.Read(buffer, binary.BigEndian, &num)
	return num
}

func BytesToInt16(buf []byte) int16  {
	var num int16
	buffer := bytes.NewBuffer(buf)
	_ = binary.Read(buffer, binary.BigEndian, &num)
	return num
}
