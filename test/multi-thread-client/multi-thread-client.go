package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	var buf [512]byte
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	service := os.Args[1]

	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err = conn.Write([]byte("hi"))
	checkError(err)

	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[0:n]))

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}
