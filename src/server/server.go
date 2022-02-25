package server

import (
	"fmt"
	"net"
	"unsafe"
)
//go-tcpsock/server.go
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		b := make([]byte, 1024)
		length,_:= c.Read(b)
		fmt.Println(string(b[0:length]))
		//fmt.Println(bytes2str(b))
		// ... ...
		// write to the connection
		//... ...
		c.Write(str2bytes("\n"))
	}
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func Start() {
	l, err := net.Listen("tcp", ":3446")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}