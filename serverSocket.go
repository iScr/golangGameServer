package main

/**
* 这是个简易的游戏服务器
* 每条包结构会有个包头标志剩下的包体长度
* 会进行断包粘包处理.
 */

import (
	"./client"
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	Head = 4
)

var (
	ClientMap map[int]net.Conn      = make(map[int]net.Conn)
	Clients   map[int]client.Client = make(map[int]client.Client)
)

func main() {
	fmt.Println(os.Args[0])
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:7981")
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	clientIndex := 0

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("监听错误: ", err.Error())
			continue
		}
		clientIndex++
		go handleClient(conn, clientIndex)
	}
}

func handleClient(conn net.Conn, index int) {
	c := client.Client{conn}
	Clients[index] = c
	ClientMap[index] = conn
	fmt.Println("新用户连接, 来自: ", conn.RemoteAddr(), "index: ", index)
	sendMsgToAll("new user added, index: " + strconv.Itoa(index))
	isHeadLoaded := false
	bodyLen := 0
	reader := bufio.NewReader(conn)

	defer func() {
		conn.Close()
		delete(ClientMap, index)
		fmt.Println("移除序号为: ", index, "的客户端")
	}()

Out:
	for {
		if !isHeadLoaded {
			headLenSl := make([]byte, 4)
			fmt.Println("读取包头....")

			//已经读取的包头字节数
			readedHeadLen := 0

			for readedHeadLen < 4 {
				len, err := reader.Read(headLenSl)
				if err != nil {
					fmt.Println("读取包头出错, ", err.Error())
					break Out
				}
				readedHeadLen += len
			}

			bodyLen = int(binary.BigEndian.Uint32(headLenSl))
			fmt.Println("读取包头成功, 包体字节长度: ", bodyLen)
			isHeadLoaded = true
		}

		if isHeadLoaded {
			fmt.Println("解析包体")
			bodySl := make([]byte, bodyLen)

			//已经读取的包体字节数
			readedBodyLen := 0

			for readedBodyLen < bodyLen {
				len, err := reader.Read(bodySl)
				if err != nil {
					fmt.Println("读取包体出错,: ", err.Error())
					break Out
				}
				readedBodyLen += len
			}

			fmt.Println("读取包体完成,包体字节长度: ", readedBodyLen)
			isHeadLoaded = false
		}
	}
}

func sendMsgToAll(msg string) {
	for _, value := range ClientMap {
		writer := bufio.NewWriter(value)
		msgLen := len(msg)
		//写入2个字节字符串长度.以供flash读取便利
		buf := make([]byte, 2)
		binary.BigEndian.PutUint16(buf, uint16(msgLen))
		writer.Write(buf)
		writer.WriteString(msg)
		writer.Flush()
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
