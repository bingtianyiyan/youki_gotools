/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/adler32"
	"io"
	"log"
	"math"
	"net"
	"testing"
)

func TestIoEncodeDemo(t *testing.T) {
	conn, _ := net.Dial("tcp", "127.0.0.1:3306")
	EncodePacket(conn, []byte("hello"))
}

type Packet struct {
	TotalSize uint32
	Magic     [4]byte
	Payload   []byte
	Checksum  uint32
}

var RPC_MAGIC = [4]byte{'p', 'y', 'x', 'i'}

func EncodePacket(w io.Writer, payload []byte) error {
	// len(Magic) + len(Checksum) == 8
	totalsize := uint32(len(RPC_MAGIC) + len(payload) + 4)
	// write total size
	binary.Write(w, binary.BigEndian, totalsize)

	sum := adler32.New()
	ww := io.MultiWriter(sum, w)
	// write magic bytes
	binary.Write(ww, binary.BigEndian, RPC_MAGIC)

	// write payload
	ww.Write(payload)

	// calculate checksum
	checksum := sum.Sum32()

	// write checksum
	return binary.Write(w, binary.BigEndian, checksum)
}

func DecodePacket(r io.Reader) ([]byte, error) {
	var totalsize uint32
	err := binary.Read(r, binary.BigEndian, &totalsize)
	if err != nil {
		return nil, errors.New("read total size")
	}

	// at least len(magic) + len(checksum)
	if totalsize < 8 {
		return nil, errors.New("bad packet. header:")
	}

	sum := adler32.New()
	rr := io.TeeReader(r, sum)

	var magic [4]byte
	err = binary.Read(rr, binary.BigEndian, &magic)
	if err != nil {
		return nil, errors.New("read magic")
	}
	if magic != RPC_MAGIC {
		return nil, errors.New("bad rpc magic")
	}


	payload := make([]byte, totalsize-8)
	_, err = io.ReadFull(rr, payload)
	if err != nil {
		return nil, errors.New("read payload")
	}

	var checksum uint32
	err = binary.Read(r, binary.BigEndian, &checksum)
	if err != nil {
		return nil, errors.New("read checksum")
	}

	if checksum != sum.Sum32() {
		return nil, errors.New(fmt.Sprintf("checkSum error, %d(calc) %d(remote)", sum.Sum32(), checksum))
	}
	return payload, nil
}

func TestIoBase64ToString(t *testing.T) {
	input := []byte("foo\x00bar")
	buffer := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)
	encoder.Write(input)
	fmt.Println(string(buffer.Bytes()))
}

func TestIoBinaryRead(t *testing.T) {
	var pi float64
	b := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
	buf := bytes.NewBuffer(b)
	err := binary.Read(buf, binary.LittleEndian, &pi)
	if err != nil {
		log.Fatalln("binary.Read failed:", err)
	}
	fmt.Println(pi)
}

func TestIoBinaryWrite(t *testing.T) {
	buf := new(bytes.Buffer)
	pi := math.Pi

	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.Bytes())
}
