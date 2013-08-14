package main

import (
	// "encoding/binary"
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
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:798")
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
	// var buffer []byte
	buffer := make([]byte, 1024)
	isHeadLoaded := false
	bodyLen := 0
	cmd := 0
	readLen := 0
	for {
		length, err := conn.Read(buffer)
		readLen += length
		checkError(err)
		fmt.Println(readLen)
		fmt.Println(bodyLen)
		fmt.Println(cmd)
		conn.Write([]byte("return"))

		if !isHeadLoaded {
			if readLen >= Head {
				fmt.Println("收到数据")
				// lenSlice := buffer[0:4]
				// bodyLen = int(binary.BigEndian.Uint32(lenSlice)) - Head
				// cmd = int(binary.BigEndian.Uint32(buffer[4:7]))

				// fmt.Println("包体长度 %d", bodyLen)
				// fmt.Println("cmd: %d", cmd)
			}
		}

	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
