package main

/*
the header model
Start (all 1s) - 1 byte
Reps (number of segments) - 4 bytes
Lengthofname - 4 bytes
Name - `lengthofname` bytes
End (all 0s) - 1 byte;

the segment model
Start (all 0s) - 1 byte
Segment number - 4 bytes
Lengthofdata - 4 bytes
Data - `lengthofdata` bytes
End (all 1s) - 1 byte

*/

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	HOST     = "localhost"
	PORT     = "8000"
	TYPE     = "tcp"
	FILENAME = "image.png"
)

type FileMetaData struct {
	name     string
	fileSize uint32
	reps     uint32
}

func prepareFIleMetaData(file *os.File) FileMetaData {
	fileInfo, err := file.Stat()

	check(err)

	size := fileInfo.Size()

	header := FileMetaData{
		name:     fileInfo.Name(),
		fileSize: uint32(size),
		reps:     uint32(size/1014) + 1,
	}

	return header
}

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	check(err)

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	sendFIle(FILENAME, conn)

	received := make([]byte, 1024)

	_, errr := conn.Read(received)

	check(errr)

	println(received)

}

func sendFIle(path string, conn *net.TCPConn) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)

	check(err)

	header := prepareFIleMetaData(file)

	dataBuffer := make([]byte, 1014)

	headerBuffer := []byte{1}

	segmentBuffer := []byte{0}

	temp := make([]byte, 4) // Temporary buffer for uint32

	received := make([]byte, 100)

	for i := 0; i < int(header.reps); i++ {
		n, _ := file.ReadAt(dataBuffer, int64(i*1014))

		if i == 0 { //send the header in the first request

			//number of segments
			binary.BigEndian.PutUint32(temp, header.reps)
			headerBuffer = append(headerBuffer, temp...)

			//length of name
			binary.BigEndian.PutUint32(temp, uint32(len(header.name)))
			headerBuffer = append(headerBuffer, temp...)

			//name of file
			headerBuffer = append(headerBuffer, []byte(header.name)...)

			headerBuffer = append(headerBuffer, 0)

			_, err := conn.Write(headerBuffer)

			check(err)

			_, err = conn.Read(received)

			check(err)

			println(string(received))

		}

		//rep number
		binary.BigEndian.PutUint32(temp, uint32(i))
		segmentBuffer = append(segmentBuffer, temp...)

		//length of data
		binary.BigEndian.PutUint32(temp, uint32(n))
		segmentBuffer = append(segmentBuffer, temp...)

		//data
		segmentBuffer = append(segmentBuffer, dataBuffer...)

		segmentBuffer = append(segmentBuffer, 1)

		_, err = conn.Write(segmentBuffer)

		check(err)

		_, err = conn.Read(received)
		fmt.Println(string(received))

		check(err)

		//reset segment buffer
		segmentBuffer = []byte{0}

	}

}
