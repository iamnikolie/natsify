package natsify

type Response struct {
	Code    int                    `json:"code"`
	Headers map[string]interface{} `json:"headers,omitempty"`
	Data    interface{}            `json:"data,omitempty"`
	Bytes   []byte                 `json:"bytes"`
}
