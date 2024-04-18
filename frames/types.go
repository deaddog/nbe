package frames

// AppId identifies the client application
type AppId string

// Serial identifies the burner server
type Serial int

type EncryptionMode string

type Function string

// MessageId identifies the message sent
// This is useful along with AppId and Serial to correlate request and response
type MessageId byte

type Password string

type ResponseCode string

type ResponsePayload map[string]string

// ControlStart is a control value placed between message header and content
const ControlStart byte = 2

// ControlEnd is a control value terminating each message
const ControlEnd byte = 4
