package main

/**
* 这是个简易的游戏服务器
* 每条包结构会有个包头标志剩下的包体长度
* 会进行断包粘包处理.
 */

import (
	// "encoding/binary"
	"fmt"
	"net"
	"os"
)

const (
	Head = 4
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
	buffer := make([]byte, 1024)
	isHeadLoaded := false
	bodyLen := 0

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(bodyLen)
		fmt.Println(length)

		if !isHeadLoaded {

			// fmt.Println("收到数据")
			// lenSlice := buffer[0:4]
			// bodyLen = int(binary.BigEndian.Uint32(lenSlice)) - Head

			// fmt.Println("包体长度 %d", bodyLen)
			// isHeadLoaded = true
		}
	}
	if isHeadLoaded {

	}

}

type myBuf []byte

func (m *myBuf) bytesAvailable() uint {
	return len(m)
}

func parseData(data []byte) {

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
