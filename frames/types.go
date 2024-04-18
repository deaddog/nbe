package frames

// AppId identifies the client application
type AppId string

// Serial identifies the burner server
type Serial int

type EncryptionMode string

const (
	EncryptionModeNone EncryptionMode = " "
	EncryptionModeRSA  EncryptionMode = "*"
	EncryptionModeXTEA EncryptionMode = "-"
)

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

// MessageId identifies the message sent
// This is useful along with AppId and Serial to correlate request and response
type MessageId byte

type Password string

type ResponseCode string

const (
	ResponseCodeOk ResponseCode = "0"
	ResponseCode1  ResponseCode = "1"
	ResponseCode2  ResponseCode = "2"
	ResponseCode3  ResponseCode = "3"
)

type ResponsePayload map[string]string

// ControlStart is a control value placed between message header and content
const ControlStart byte = 2

// ControlEnd is a control value terminating each message
const ControlEnd byte = 4
