package main

import (
	. "fmt"
	"net"
	"os"
)

func main() {
	service := "0.0.0.0:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleclient(conn)
	}
}
func handleclient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		Println("go client Connect")
		Println(string(buf[0:n]))
		_, err = conn.Write(buf[0:n])
		if err != nil {
			return
		}
	}
}

func checkError(err error) {
	if err != nil {
		Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
