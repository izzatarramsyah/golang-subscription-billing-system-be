package models

type Response struct {
    Status  int         `json:"status"`  // HTTP Status code
    Message string      `json:"message"` // Pesan yang terkait dengan hasil request
    Data    interface{} `json:"data"`    // Data hasil response
}

func NewSuccessResponse(data interface{}, message string) Response {
    return Response{
        Status:  200,
        Message: message,
        Data:    data,
    }
}

func NewErrorResponse(status int, message string) Response {
    return Response{
        Status:  status,
        Message: message,
        Data:  nil,
    }
}