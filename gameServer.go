package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)

		// fmt.Println(conn.RemoteAddr())
		// fmt.Println(conn.LocalAddr())

		// daytime := time.Now().String()
		// fmt.Println("connected")
		// conn.Write([]byte(daytime))
		// conn.Close()
	}
}

func handleClient(conn net.Conn) {
	fmt.Println(conn.RemoteAddr())
	dayTime := time.Now().String()
	fmt.Println(dayTime)

}

func checkError(err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
