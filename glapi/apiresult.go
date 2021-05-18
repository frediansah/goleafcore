package glapi

type OutgoingApi struct {
	Result OutgoingApiResult `json:"result"`
}

type OutgoingApiResult struct {
	Status       string        `json:"status"`
	ErrorCode    string        `json:"errorCode,omitempty"`
	ErrorMsg     string        `json:"errorMsg,omitempty"`
	ErrorFullMsg string        `json:"errorFullMsg,omitempty"`
	ErrorArgs    []interface{} `json:"errorArgs,omitempty"`

	Payload interface{} `json:"payload,omitempty"`
}
