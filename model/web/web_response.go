package web

type BaseResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ResponseWithData struct {
	BaseResponse
	Data interface{} `json:"data"`
}

type ResponseWithMesage struct {
	BaseResponse
	Message string `json:"message"`
}

type Response struct {
	BaseResponse
	Data interface{} `json:"data"`
	Message string `json:"message"`
}

func NewResponseWithData(code int, status string, data interface{}) ResponseWithData {
	return ResponseWithData{
		BaseResponse: BaseResponse{
			Code:   code,
			Status: status,
		},
		Data:    data,
	}
}

func NewResponseWithMessage(code int, status string, message string) ResponseWithMesage {
	return ResponseWithMesage{
		BaseResponse: BaseResponse{
			Code:   code,
			Status: status,
		},
		Message: message,
	}
}

func NewResponse(code int, status string, data interface{}, message string) Response {
	return Response{
		BaseResponse: BaseResponse{
			Code:   code,
			Status: status,
		},
		Data:    data,
		Message: message,
	}
}