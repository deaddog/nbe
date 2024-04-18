package frames

type Response struct {
	AppId    AppId
	Serial   Serial
	Function Function
	Id       MessageId
	Code     ResponseCode
	Payload  ResponsePayload
}
