package main
import (
	"fmt"
	"net"
)

func main(){
	ln, err := net.Listen("tcp", ":5555")
	if err != nil{
		fmt.Println(err)
		return
	}

	for{
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){

	buf := make([]byte, 6000)
	_, err := conn.Read(buf)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("Received: " + string(buf))
}