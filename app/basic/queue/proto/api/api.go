package api

type Request struct {
	Name string `path:"name,options=you|me"`
}

type HelloReq struct {
	Name string `path:"name"`
}

type Response struct {
	Message string `json:"message"`
}

// queue message:
type Message struct {
	Type      string                 `json:"type"`
	Data      map[string]interface{} `json:"data"`
	Timestamp int                    `json:"timestamp"`
}
