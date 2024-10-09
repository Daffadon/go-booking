package utils

type ResponseError struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func ReturnResponseError(statusCode int, message string) ResponseError {
	return ResponseError{
		StatusCode: statusCode,
		Message:    message,
	}
}

func ReturnResponseSuccess(statusCode int, message string, data interface{}) ResponseError {
	return ResponseError{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
