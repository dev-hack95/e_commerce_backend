package utilities

type ResponseJson struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Model interface{} `json:"model"`
}

func SuccessResponse(returnData *ResponseJson, data interface{}) {
	returnData.Code = 200
	returnData.Msg = "success"
	returnData.Model = data
}

func ErrorResponse(returnData *ResponseJson, msg string) {
	returnData.Code = 400
	returnData.Msg = msg
}

func UnAuthorizationResponse(returnData *ResponseJson) {
	returnData.Code = 404
	returnData.Msg = "Authorization Error"
	returnData.Model = nil
}
