package main

/**
* 这是个简易的游戏服务器
* 每条包结构会有个包头标志剩下的包体长度
* 会进行断包粘包处理.
 */

import (
	"bufio"
	"encoding/binary"
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
	isHeadLoaded := false
	bodyLen := 0
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		// buffer := make([]byte, 1024)
		// length, _ := conn.Read(buffer)
		// fmt.Println(length)
		// if length < 8 {
		// 	continue
		// }
		// if err != nil {
		// 	fmt.Println(err)
		// 	break
		// }
		// fmt.Println(bodyLen)
		// fmt.Println(length)

		if !isHeadLoaded {
			lenSl := make([]byte, 4)
			fmt.Println("读取包头....")
			len, err := reader.Read(lenSl)
			if err != nil {
				fmt.Println("读取包头出错, ", err.Error())
				break
			}
			fmt.Println("读取到的包头长度: ", len)

			bodyLen = int(binary.BigEndian.Uint32(lenSl))
			fmt.Println("包体字节长度: ", bodyLen)
			isHeadLoaded = true
			// fmt.Println("收到数据")
			// lenSlice := buffer[0:4]
			// bodyLen = int(binary.BigEndian.Uint32(lenSlice)) - Head

			// fmt.Println("包体长度 %d", bodyLen)
			// isHeadLoaded = true
		}
		if isHeadLoaded {
			fmt.Println("解析包体")
			bodySl := make([]byte, bodyLen)
			len, err := reader.Read(bodySl)
			if err != nil {
				fmt.Println("读取包体出错,: ", err.Error())
				break
			}
			fmt.Println("读取到的包体长度: ", len)

			isHeadLoaded = false

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
