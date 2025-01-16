package enums

type ExampleCodeSandbox struct {
}

func (e *ExampleCodeSandbox) ExecuteCode(request ExecuteCodeRequest) ExecuteCodeResponse {
	inputList := request.InputList
	executeCodeResponse := new(ExecuteCodeResponse)
	executeCodeResponse.OutputList = inputList
	executeCodeResponse.Message = "测试执行成功"
	executeCodeResponse.Status = "SUCCESS"
	judgeInfo := new(JudgeInfo)
	judgeInfo.Message = "SUCCESS"
	judgeInfo.Memory = 100
	judgeInfo.Time = 100
	executeCodeResponse.JudgeInfo = judgeInfo
	return *executeCodeResponse
}
