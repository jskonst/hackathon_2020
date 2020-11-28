package main

import (
	"testing"
	"time"
)

func TestGetMessageLocation(t *testing.T) {
	expMsg := Message{
		msgtype:  "BR00",
		deviceId: "072101913184",
		payload:  "201127V5659.3750N04058.6176E002.5083805192.29,00000000L00000000",
	}
	msg := GetMessage("(072101913184BR00201127V5659.3750N04058.6176E002.5083805192.29,00000000L00000000)")
	if expMsg != msg {
		t.Error("For", "Message", "expected", expMsg, "got", msg)
	}
}

func TestGetMessageHandshake(t *testing.T) {
	expMsg := Message{
		msgtype:  "BP00",
		deviceId: "072101913184",
		payload:  "352672101913184HSOP51",
	}
	msg := GetMessage("(072101913184BP00352672101913184HSOP51)")
	if expMsg != msg {
		t.Error("For", "Message", "expected", expMsg, "got", msg)
	}
}

func TestConvertCoordinate1(t *testing.T) {
	exp := 56.989583
	res := convertPosition("5659.3750", 2)
	if (res - exp) > 0.000001 {
		t.Error("For coordinate conversion 5659.3750", "expected", exp, "got", res)
	}
}

func TestConvertCoordinate2(t *testing.T) {
	exp := 40.97696
	res := convertPosition("04058.6176", 3)
	if (res - exp) > 0.000001 {
		t.Error("For coordinate conversion 5659.3750", "expected", exp, "got", res)
	}
}

func TestConvertDate(t *testing.T) {
	exp := time.Date(2020, time.November, 27, 8, 38, 05, 0, time.UTC)
	res := convertDate("201127", "083805")
	if res != exp {
		t.Error("For date conversion", "expected", exp, "got", res)
	}
}

func TestGetPosition(t *testing.T) {
	msg := Message{
		msgtype:  "BR00",
		deviceId: "072101913184",
		payload:  "201127V5659.3750N04058.6176E002.5083805192.29,00000000L00000000",
	}
	exp := Position{
		deviceId:     "072101913184",
		latitude:     56.989583,
		longitude:    40.97696,
		speed:        2.5,
		orientation:  192.29,
		isAvaliable:  false,
		latIndicator: "N",
		lonIndicator: "E",
		dateTime:     time.Date(2020, time.November, 27, 8, 38, 05, 0, time.UTC),
	}
	res := GetPosition(msg)
	if res.deviceId != exp.deviceId {
		t.Error("For deviceId", "expected", exp.deviceId, "got", res.deviceId)
	}
	if res.latitude-exp.latitude > 0.00001 {
		t.Error("For deviceId", "expected", exp.latitude, "got", res.latitude)
	}
	if res.longitude-exp.longitude > 0.00001 {
		t.Error("For deviceId", "expected", exp.longitude, "got", res.longitude)
	}
	if res.speed-exp.speed > 0.01 {
		t.Error("For deviceId", "expected", exp.speed, "got", res.speed)
	}
	if res.orientation-exp.orientation > 0.01 {
		t.Error("For deviceId", "expected", exp.orientation, "got", res.orientation)
	}
	if res.isAvaliable != exp.isAvaliable {
		t.Error("For deviceId", "expected", exp.isAvaliable, "got", res.isAvaliable)
	}
	if res.latIndicator != exp.latIndicator {
		t.Error("For deviceId", "expected", exp.latIndicator, "got", res.latIndicator)
	}
	if res.lonIndicator != exp.lonIndicator {
		t.Error("For deviceId", "expected", exp.lonIndicator, "got", res.lonIndicator)
	}
	if res.dateTime != exp.dateTime {
		t.Error("For deviceId", "expected", exp.dateTime, "got", res.dateTime)
	}

}

func TestGenerateHandshake(t *testing.T) {
	incMsg := Message{
		msgtype:  "BP00",
		deviceId: "072101913184",
		payload:  "352672101913184HSOP51",
	}
	msg := HandshakeResponse(incMsg)
	if msg != "(072101913184AP01HSO)" {
		t.Error("For", "Handshake", "expected", "(072101913184AP01HSO)", "got", msg)
	}
}
