package enums

// CodeSandbox 定义一个接口，所有沙箱类型都需要实现它
type CodeSandbox interface {
	ExecuteCode(request ExecuteCodeRequest) ExecuteCodeResponse
}

// ExecuteCodeRequest 请求结构体
type ExecuteCodeRequest struct {
	Code      string
	Language  string
	InputList []string
}

// ExecuteCodeResponse 响应结构体
type ExecuteCodeResponse struct {
	OutputList []string
	Message    string
	Status     string
	JudgeInfo  *JudgeInfo
}

// JudgeInfo 判题信息
type JudgeInfo struct {
	Message string
	Memory  int
	Time    int
}
