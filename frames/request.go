package frames

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	AppId          AppId
	Serial         Serial
	EncryptionMode EncryptionMode
	Function       Function
	Id             MessageId
	Password       Password
	Timestamp      time.Time
	Payload        string
}

// The "pad " string is 4 bytes reserved for "future" use.
// They are not used in any messages, but validated to ensure message correctness.
const fieldExtra = "pad "

func EncodeRequest(r Request) []byte {
	data := []byte{}

	payloadsize := fmt.Sprintf("%03d", len(r.Payload))

	data = append(data, []byte(fmt.Sprintf("%12s", r.AppId))...)
	data = append(data, []byte(fmt.Sprintf("%06d", r.Serial))...)
	data = append(data, []byte(r.EncryptionMode)...)
	data = append(data, ControlStart)
	data = append(data, []byte(r.Function)...)
	data = append(data, []byte(fmt.Sprintf("%02d", r.Id))...)
	data = append(data, []byte(r.Password)...)
	data = append(data, []byte(strconv.FormatInt(r.Timestamp.Unix(), 10))...)
	data = append(data, []byte(fieldExtra)...)
	data = append(data, []byte(payloadsize)...)
	data = append(data, []byte(r.Payload)...)
	data = append(data, ControlEnd)

	return data
}

func DecodeRequest(buffer []byte) (Request, error) {
	size := string(buffer[48:51])
	actualSize, err := strconv.Atoi(size)
	if err != nil {
		return Request{}, fmt.Errorf("could not parse payload size: %w", err)
	}

	id, err := strconv.Atoi(string(buffer[22:24]))
	if err != nil {
		return Request{}, fmt.Errorf("could not parse message id '%s': %w", string(buffer[22:24]), err)
	}

	tsUnix, err := strconv.ParseInt(string(buffer[34:44]), 10, 64)
	if err != nil {
		return Request{}, fmt.Errorf("could not parse unix timestamp '%s': %w", string(buffer[34:44]), err)
	}
	ts := time.Unix(tsUnix, 0)

	if buffer[19] != ControlStart {
		return Request{}, fmt.Errorf("start '%d' must be %d", buffer[19], ControlStart)
	}

	if string(buffer[44:48]) != fieldExtra {
		return Request{}, fmt.Errorf("extra '%s' must be exacly '%s'", string(buffer[44:48]), fieldExtra)
	}

	if buffer[51+actualSize] != ControlEnd {
		return Request{}, fmt.Errorf("end '%d' must be %d", buffer[51+actualSize], ControlEnd)
	}

	serial, err := strconv.Atoi(string(buffer[12:18]))
	if err != nil {
		return Request{}, fmt.Errorf("err reading burner serial: %w", err)
	}

	return Request{
		AppId:          AppId(strings.Trim(string(buffer[0:12]), " ")),
		Serial:         Serial(serial),
		EncryptionMode: EncryptionMode(buffer[18]),
		Function:       Function(buffer[20:22]),
		Id:             MessageId(id),
		Password:       Password(buffer[24:34]),
		Timestamp:      ts,
		Payload:        string(buffer[51 : 51+actualSize]),
	}, nil
}
