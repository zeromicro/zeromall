package api

type HelloReq struct {
	Name string `path:"name"`
}

// queue message:
type MessageReq struct {
	QueueName string `json:"queue_name"`

	Type string `json:"type"`
	Data string `json:"data"`
	//RawData   []byte `json:"raw_data"`
	Timestamp int `json:"timestamp"`
}

//
type MessageResp struct {
	Status string `json:"status"`
}
