package protocol

import "encoding/json"

type ResponseCode string

const (
	ResponseCode_Success ResponseCode = "success"
	ResponseCode_Failure ResponseCode = "failure"
)

type Response struct {
	Code    ResponseCode    `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}
