package frames

import "fmt"

// AppId identifies the client application
type AppId string

func (a AppId) Validate() error {
	if len(a) > 12 {
		return fmt.Errorf("app id '%s' must be no longer than 12 characters", a)
	}
	if len(a) < 1 {
		return fmt.Errorf("app id cannot be empty")
	}
	return nil
}

// Serial identifies the burner server
type Serial int

func (s Serial) Validate() error {
	if s < 0 {
		return fmt.Errorf("serial '%v' must be positive", s)
	}
	if s > 999999 {
		return fmt.Errorf("serial '%v' must be at most 6 digits", s)
	}
	return nil
}

type EncryptionMode string

const (
	EncryptionModeNone EncryptionMode = " "
	EncryptionModeRSA  EncryptionMode = "*"
	EncryptionModeXTEA EncryptionMode = "-"
)

func (e EncryptionMode) Validate() error {
	if e != EncryptionModeNone && e != EncryptionModeRSA && e != EncryptionModeXTEA {
		return fmt.Errorf("encryption mode '%s' must be '%s'(%s), '%s'(%s), or '%s'(%s)", e, EncryptionModeNone, "none", EncryptionModeRSA, "RSA", EncryptionModeXTEA, "xtea")
	}
	return nil
}

type Function string

const (
	FunctionDiscovery             Function = "00"
	FunctionReadSetupValue        Function = "01"
	FunctionSetSetupValue         Function = "02"
	FunctionReadSetupRange        Function = "03"
	FunctionReadOperatingData     Function = "04"
	FunctionReadAdvancedData      Function = "05"
	FunctionReadConsumptionData   Function = "06"
	FunctionReadChartData         Function = "07"
	FunctionReadEventLog          Function = "08"
	FunctionReadInfo              Function = "09"
	FunctionReadAvailablePrograms Function = "10"
)

func (f Function) Validate() error {
	if f != FunctionDiscovery &&
		f != FunctionReadSetupValue &&
		f != FunctionSetSetupValue &&
		f != FunctionReadSetupRange &&
		f != FunctionReadOperatingData &&
		f != FunctionReadAdvancedData &&
		f != FunctionReadConsumptionData &&
		f != FunctionReadChartData &&
		f != FunctionReadEventLog &&
		f != FunctionReadInfo &&
		f != FunctionReadAvailablePrograms {
		return fmt.Errorf("unknown function '%s'", f)
	}
	return nil
}

// MessageId identifies the message sent
// This is useful along with AppId and Serial to correlate request and response
type MessageId byte

func (id MessageId) Validate() error {
	if id >= 100 {
		return fmt.Errorf("message id '%d' must be in range [0; 100], both inclusive", id)
	}
	return nil
}

type Password string

func (p Password) Validate() error {
	if len(p) != 10 && len(p) != 0 {
		return fmt.Errorf("password '%s' must be 10 characters", p)
	}
	return nil
}

type ResponseCode string

const (
	ResponseCodeOk ResponseCode = "0"
	ResponseCode1  ResponseCode = "1"
	ResponseCode2  ResponseCode = "2"
	ResponseCode3  ResponseCode = "3"
)

func (r ResponseCode) Validate() error {
	if r != ResponseCodeOk &&
		r != ResponseCode1 &&
		r != ResponseCode2 &&
		r != ResponseCode3 {
		return fmt.Errorf("response code '%s' must be one of %s, %s, %s, %s", r, ResponseCodeOk, ResponseCode1, ResponseCode2, ResponseCode3)
	}

	return nil
}

type ResponsePayload map[string]string

// ControlStart is a control value placed between message header and content
const ControlStart byte = 2

// ControlEnd is a control value terminating each message
const ControlEnd byte = 4
