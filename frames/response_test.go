package frames_test

import (
	"reflect"
	"testing"

	"github.com/deaddog/nbe/frames"
)

func TestResponseEncoding(t *testing.T) {
	original := frames.Response{
		AppId:    "abc",
		Serial:   12345,
		Function: frames.FunctionReadConsumptionData,
		Id:       12,
		Code:     frames.ResponseCodeOk,
		Payload: map[string]string{
			"field": "value",
		},
	}

	encoded := frames.EncodeResponse(original)
	decoded, err := frames.DecodeResponse(encoded)

	if err != nil {
		t.Fatalf("decoding failed: %v", err)
	}

	if original.AppId != decoded.AppId {
		t.Fatalf("AppId '%v' was '%v' efter encode->decode", original.AppId, decoded.AppId)
	}
	if original.Serial != decoded.Serial {
		t.Fatalf("Serial '%v' was '%v' efter encode->decode", original.Serial, decoded.Serial)
	}
	if original.Function != decoded.Function {
		t.Fatalf("Function '%v' was '%v' efter encode->decode", original.Function, decoded.Function)
	}
	if original.Id != decoded.Id {
		t.Fatalf("Id '%v' was '%v' efter encode->decode", original.Id, decoded.Id)
	}
	if original.Code != decoded.Code {
		t.Fatalf("Code '%v' was '%v' efter encode->decode", original.Code, decoded.Code)
	}
	if !reflect.DeepEqual(original, decoded) {
		t.Fatalf("Payload '%v' was '%v' efter encode->decode", original.Payload, decoded.Payload)
	}
}
