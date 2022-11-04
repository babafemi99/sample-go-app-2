package models

type ArithmeticReq struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"validate:"required"`
	Y             int    `json:"y"validate:"required"`
}

type ArithmeticRes struct {
	SlackUsername string `json:"slackUsername"`
	Result        int    `json:"result"`
	OperationType string `json:"operation_type"`
}
