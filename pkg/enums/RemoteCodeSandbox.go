package enums

import (
	"fmt"
)

type RemoteCodeSandbox struct {
}

func (e *RemoteCodeSandbox) ExecuteCode(request ExecuteCodeRequest) ExecuteCodeResponse {
	fmt.Println("Executing code in RemoteCodeSandbox...")
	// 这里可以执行实际的代码
	outputList := []string{
		"Output Line 1",
		"Output Line 2",
	}

	// 模拟判题信息
	judgeInfo := JudgeInfo{
		Message: "Code executed successfully",
		Memory:  256, // 假设内存使用为 256 MB
		Time:    120, // 假设代码执行时间为 120ms
	}

	// 返回响应
	return ExecuteCodeResponse{
		OutputList: outputList,           // 返回的输出列表
		Message:    "Execution complete", // 返回的消息
		Status:     "ERROR",
		JudgeInfo:  &judgeInfo, // 填充判题信息
	}
}
