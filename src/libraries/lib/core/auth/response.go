package auth

type Errors struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Field   string `json:"field,omitempty"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  Errors      `json:"errors,omitempty"` //TODO : make it in array JSON
}

func (res *Response) Set(status string, message string, data interface{}, err Errors) {
	res.Status = status
	res.Message = message

	res.Errors = err

	if data != nil {
		res.Data = data
	}
}
