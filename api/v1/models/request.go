package models

type Request struct {
	Data map[string]interface{} `json:"data"`
}

// NewGeneralRequest digunakan untuk membuat request body dengan data yang lebih fleksibel
func NewGeneralRequest(data map[string]interface{}) Request {
    return Request{
        Data: data,
    }
}