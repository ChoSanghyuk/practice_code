package parameters

type CommonResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) CommonResponse {
	return CommonResponse{
		Data: data,
	}
}

func NewErrorResponse(errMsg string) CommonResponse {
	return CommonResponse{
		Message: errMsg,
	}
}
