package util

const (
	ERROR   = 0
	SUCCESS = 1
	FAILED  = 9
)

// 响应对象
type ReturnObject struct {
	ReturnType    int         // 状态
	ReturnMessage string      // 消息
	ReturnData    interface{} // 数据
}

// 返回错误对象
func Error(err error) *ReturnObject {
	objReturn := new(ReturnObject)
	objReturn.ReturnType = ERROR
	objReturn.ReturnMessage = err.Error()
	objReturn.ReturnData = err
	return objReturn
}

// 返回成功对象
func Success(data interface{}, msgs ...string) *ReturnObject {
	objReturn := new(ReturnObject)
	objReturn.ReturnData = data
	objReturn.ReturnType = SUCCESS
	msg := "Success"
	if len(msgs) > 0 {
		msg = msgs[0]
	}
	objReturn.ReturnMessage = msg
	return objReturn
}

// 返回失败对象
func Failed(data interface{}, msgs ...string) *ReturnObject {
	objReturn := new(ReturnObject)
	objReturn.ReturnData = data
	objReturn.ReturnType = FAILED
	if len(msgs) > 0 {
		objReturn.ReturnMessage = msgs[0]
	}

	return objReturn
}
