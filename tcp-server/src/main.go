package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// buf := make([]byte, 1024)
	listner, err := net.Listen("tcp", ":6000")
	if err != nil {
		log.Fatal("Error 0 ", err)
		return
	}
	defer listner.Close()
	log.Printf("Server started")
	for {
		conn, err := listner.Accept()
		if err != nil {
			fmt.Println("Error1 ", err)
			return
		}

		r := bufio.NewReader(conn)
		w := bufio.NewWriter(conn)

		line, err := r.ReadString(')')
		if err != nil {
			fmt.Println("Error 2 ", err)
			conn.Close()
			continue
		}
		fmt.Println(line)
		message := GetMessage(line)
		switch message.msgtype {
		case "BP00":
			resp := HandshakeResponse(message)
			w.WriteString(resp)
			fmt.Println("Wrote handshake resp")
			break
		case "BR00":
			result := GetPosition(message)
			fmt.Println("Got position")
			fmt.Println(result)
			fmt.Println(result.dateTime)
			break

		}
		w.Flush()

		conn.Close()
	}
}
