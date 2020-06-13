package response

// APIResponse :
type APIResponse struct {
	Code    int         `json:"code" binding:"required"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

/*Nil :
 */
var (
	DataListNil = make([]int, 0)           // []
	DataItemNil = map[string]interface{}{} // {}
)

// APIResponseData :
func APIResponseData(code int, message string, data interface{}) *APIResponse {
	res := new(APIResponse)
	res.Code = code
	if message == "" {
		res.Message = "OK"
	} else {
		res.Message = message
	}
	if data != nil {
		res.Data = data
	} else {
		res.Data = DataItemNil //map[string]interface{}{}
	}
	return res
}

// Pagination :分页属性
type Pagination struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"pageSize"`
}
