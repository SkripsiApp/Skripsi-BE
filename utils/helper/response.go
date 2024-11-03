package helper

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type ErrorResponseMiddleware struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type SuccessResponseWithPaginationAndCount struct {
	Status     bool        `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
	Count      int         `json:"count_data,omitempty"`
}

func (e *ErrorResponse) Error() string {
	return e.Message
}

func ResponseError(statusCode int, message string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}
}

func ResponseErrorMiddleware(message string) ErrorResponseMiddleware {
	return ErrorResponseMiddleware{
		Status:  false,
		Message: message,
	}
}

func ResponseSuccess(message string) SuccessResponse {
	return SuccessResponse{
		Status:  true,
		Message: message,
	}
}

func ResponseSuccessWithData(message string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

func ResponseSuccessWithPagnationAndCount(message string, data interface{}, pagnation interface{}, count int) SuccessResponseWithPaginationAndCount {
	return SuccessResponseWithPaginationAndCount{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagnation,
		Count:      count,
	}
}