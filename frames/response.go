package frames

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Response struct {
	AppId    AppId
	Serial   Serial
	Function Function
	Id       MessageId
	Code     ResponseCode
	Payload  ResponsePayload
}

func (r *Response) Validate() error {
	err := errors.Join(
		r.AppId.Validate(),
		r.Serial.Validate(),
		r.Function.Validate(),
		r.Id.Validate(),
		r.Code.Validate(),
	)

	for k, v := range r.Payload {
		if strings.ContainsAny(k, ";=") {
			err = errors.Join(fmt.Errorf("payload key '%s' cannot contain ; and =", k))
		}
		if strings.ContainsAny(v, ";=") {
			err = errors.Join(fmt.Errorf("payload value '%s' on key '%s' cannot contain ; and =", v, k))
		}
	}

	return err
}

func EncodeResponse(r Response) []byte {
	data := []byte{}

	payload := ""
	for k, v := range r.Payload {
		payload += fmt.Sprintf(";%s=%s", k, v)
	}
	if len(payload) > 0 {
		payload = payload[1:]
	}
	payloadsize := fmt.Sprintf("%03d", len(payload))

	data = append(data, []byte(fmt.Sprintf("%12s", r.AppId))...)
	data = append(data, []byte(fmt.Sprintf("%06d", r.Serial))...)
	data = append(data, ControlStart)
	data = append(data, []byte(r.Function)...)
	data = append(data, []byte(fmt.Sprintf("%02d", r.Id))...)
	data = append(data, []byte(r.Code)...)
	data = append(data, []byte(payloadsize)...)
	data = append(data, []byte(payload)...)
	data = append(data, ControlEnd)

	return data
}

func DecodeResponse(buffer []byte) (Response, error) {
	size := string(buffer[24:27])
	actualSize, err := strconv.Atoi(size)
	if err != nil {
		return Response{}, fmt.Errorf("could not parse payload size: %w", err)
	}

	id, err := strconv.Atoi(string(buffer[21:23]))
	if err != nil {
		return Response{}, fmt.Errorf("could not parse message id '%s': %w", string(buffer[21:23]), err)
	}

	if buffer[18] != ControlStart {
		return Response{}, fmt.Errorf("start '%d' must be %d", buffer[19], ControlStart)
	}

	if buffer[27+actualSize] != ControlEnd {
		return Response{}, fmt.Errorf("end '%d' must be %d", buffer[27+actualSize], ControlEnd)
	}

	payload := map[string]string{}

	for _, pair := range strings.Split(string(buffer[27:27+actualSize]), ";") {
		vals := strings.SplitN(pair, "=", 2)
		if len(vals) != 2 {
			return Response{}, fmt.Errorf("could not parse payload key=value '%s'", pair)
		}
		payload[vals[0]] = vals[1]
	}

	serial, err := strconv.Atoi(string(buffer[12:18]))
	if err != nil {
		return Response{}, fmt.Errorf("err reading burner serial: %w", err)
	}

	return Response{
		AppId:    AppId(strings.Trim(string(buffer[0:12]), " ")),
		Serial:   Serial(serial),
		Function: Function(buffer[19:21]),
		Id:       MessageId(id),
		Code:     ResponseCode(buffer[23:24]),
		Payload:  payload,
	}, nil
}
