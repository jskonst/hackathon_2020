package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

type Message struct {
	msgtype  string
	deviceId string
	payload  string
}

type Position struct {
	deviceId     string
	latitude     float64
	longitude    float64
	speed        float64
	orientation  float64
	isAvaliable  bool
	latIndicator string
	lonIndicator string
	dateTime     time.Time
}

func GetMessage(msg string) Message {
	receivedMessage := Message{}
	pattern := regexp.MustCompile(`\((\d{12})(\w{2}\d{2})(.*)\)`)
	matched := pattern.FindSubmatch([]byte(msg))
	if len(matched) == 4 {
		receivedMessage.deviceId = string(matched[1])
		receivedMessage.msgtype = string(matched[2])
		receivedMessage.payload = string(matched[3])
	}
	return receivedMessage
}

func convertPosition(coordinate string, degNum int) float64 {
	deg, _ := strconv.ParseFloat(coordinate[:degNum], 64)
	min, _ := strconv.ParseFloat(coordinate[degNum:], 64)
	return (float64)(deg + (min / 60))
}

func convertDate(date string, tm string) time.Time {
	value := fmt.Sprintf("%s %s", date, tm)
	layout := "060102 030405"
	t, _ := time.Parse(layout, value)
	return t
}

func GetPosition(message Message) Position {
	result := Position{}
	result.deviceId = message.deviceId
	pattern := regexp.MustCompile(`(\d{6})(\D{1})([\d.]{9})(\D{1})([\d.]{10})(\D{1})([\d.]{5})(\d{6})([\d.]{6})`)
	matched := pattern.FindSubmatch([]byte(message.payload))
	if len(matched) == 10 {
		result.dateTime = convertDate(string(matched[1]), string(matched[8]))
		avaliability := string(matched[2])
		switch avaliability {
		case "A":
			result.isAvaliable = true
			break
		case "V":
			result.isAvaliable = false
		}
		result.latitude = convertPosition(string(matched[3]), 2)
		result.latIndicator = string(matched[4])
		result.longitude = convertPosition(string(matched[5]), 3)
		result.lonIndicator = string(matched[6])
		result.speed, _ = strconv.ParseFloat(string(matched[7]), 32)
		result.orientation, _ = strconv.ParseFloat(string(matched[9]), 32)
	}

	return result
}

func HandshakeResponse(msg Message) string {
	return fmt.Sprintf("(%sAP01HSO)", msg.deviceId)
}
