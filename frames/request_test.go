package frames_test

import (
	"testing"
	"time"

	"github.com/deaddog/nbe/frames"
)

func TestRequestEncoding(t *testing.T) {
	original := frames.Request{
		AppId:          "abc",
		Serial:         12345,
		EncryptionMode: frames.EncryptionModeNone,
		Function:       frames.FunctionReadConsumptionData,
		Id:             12,
		Password:       "          ",
		Timestamp:      time.Now().Truncate(time.Second),
		Payload:        "",
	}

	encoded := frames.EncodeRequest(original)
	decoded, err := frames.DecodeRequest(encoded)

	if err != nil {
		t.Fatalf("decoding failed: %v", err)
	}

	if original.AppId != decoded.AppId {
		t.Fatalf("AppId '%v' was '%v' efter encode->decode", original.AppId, decoded.AppId)
	}
	if original.Serial != decoded.Serial {
		t.Fatalf("Serial '%v' was '%v' efter encode->decode", original.Serial, decoded.Serial)
	}
	if original.EncryptionMode != decoded.EncryptionMode {
		t.Fatalf("EncryptionMode '%v' was '%v' efter encode->decode", original.EncryptionMode, decoded.EncryptionMode)
	}
	if original.Function != decoded.Function {
		t.Fatalf("Function '%v' was '%v' efter encode->decode", original.Function, decoded.Function)
	}
	if original.Id != decoded.Id {
		t.Fatalf("Id '%v' was '%v' efter encode->decode", original.Id, decoded.Id)
	}
	if original.Password != decoded.Password {
		t.Fatalf("Password '%v' was '%v' efter encode->decode", original.Password, decoded.Password)
	}
	if original.Timestamp != decoded.Timestamp {
		t.Fatalf("Timestamp '%v' was '%v' efter encode->decode", original.Timestamp, decoded.Timestamp)
	}
	if original.Payload != decoded.Payload {
		t.Fatalf("Payload '%v' was '%v' efter encode->decode", original.Payload, decoded.Payload)
	}
}
