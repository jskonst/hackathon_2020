package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jskonst/hackathon_2020/tcp-server/config"
	"log"
	"net"
	"net/http"
)

func main() {
	cfg, err := config.New("../.env")
	if err != nil {
		log.Fatal(err)
	}

	listen(cfg.ListenAddress, func(position Position) {
		log.Println("GO")
		sendAddPositionRequest(cfg.APIAddress + "/api/positions", position)
	})
}

// listen ...
func listen(listenAddress string, onNewPosition func(Position)) {
	listner, err := net.Listen("tcp", listenAddress)
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
			position := GetPosition(message)
			fmt.Println("Got position")
			fmt.Println(position)
			fmt.Println(position.dateTime)
			onNewPosition(position)
			break

		}

		w.Flush()
		conn.Close()
	}
}

// sendAddPositionRequest ...
func sendAddPositionRequest(apiAddress string, position Position) {
	var requestModel AddPositionRequestModel

	requestModel.IMEI = position.deviceId
	requestModel.Timestamp = position.dateTime
	requestModel.Latitude = position.latitude
	requestModel.Longitude = position.longitude

	requestBody, err := json.Marshal(requestModel)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post(apiAddress, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalf("API returns status code: %d", response.StatusCode)
	}
}
