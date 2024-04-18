package frames

import "time"

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
