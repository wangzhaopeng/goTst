package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func Handle_conn(conn net.Conn) {

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			conn.Close()
			log.Fatal(err)
			fmt.Printf("read err\n")
			return
		} else if err == io.EOF {
			conn.Close()
			fmt.Printf("close\n")
			return
		} else {
			time.Sleep(time.Second)
			nW, err := conn.Write(buf[:n])
			if err != nil || nW != n {
				conn.Close()
				fmt.Printf("write err\n")
				return
			}
		}
	}

}

func main() {
	addr := "0.0.0.0:8080"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept() //用conn接收链接
		if err != nil {
			log.Fatal(err)
		}
		go Handle_conn(conn) //开启多个协程。
	}
}
