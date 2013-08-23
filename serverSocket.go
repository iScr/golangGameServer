package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	Len  = 4
	Cmd  = 4
	Head = Len + Cmd
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7981")
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}

}

func handleClient(conn net.Conn) {
	fmt.Println(conn.RemoteAddr())
	// var readBuffer []byte
	readBuffer := make([]byte, 1024)
	// buffer := make([]byte, 1024)
	isHeadLoaded := false
	bodyLen := 0
	// cmd := 0

	readLen := 0

	for {

		length, err := conn.Read(readBuffer[readLen:])
		readLen += length
		checkError(err)
		fmt.Println("readLen: ", readLen)
		fmt.Println("bodyLen: ", bodyLen)
		fmt.Println(string(readBuffer[0:readLen]))
		// fmt.Println("cmd: ", cmd)
		conn.Write([]byte("return"))

		if !isHeadLoaded {
			continue
			if readLen > Head {
				fmt.Println("收到数据")
				lenSlice := readBuffer[0:4]
				bodyLen = int(binary.BigEndian.Uint32(lenSlice)) - Head
				cmd = int(binary.BigEndian.Uint32(readBuffer[4:8]))

				fmt.Println("包体长度 %d", bodyLen)
				// fmt.Println("cmd: %d", cmd)
				isHeadLoaded = true
			}
		}
		if isHeadLoaded {
			if readLen >= bodyLen {

			}

		}

	}
}

func parseData(data []byte) {

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
