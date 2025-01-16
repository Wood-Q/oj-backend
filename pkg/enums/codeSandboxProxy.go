package enums

import (
	"fmt"
)

// CodeSandboxProxy 代理类
type CodeSandboxProxy struct {
	// 包含一个 CodeSandbox 实现
	realSandbox CodeSandbox
}

// NewCodeSandboxProxy 创建一个新的代理
func NewCodeSandboxProxy(sandbox CodeSandbox) *CodeSandboxProxy {
	return &CodeSandboxProxy{realSandbox: sandbox}
}

// ExecuteCode 方法实现了 CodeSandbox 接口
// 代理在这里增加了额外操作
func (p *CodeSandboxProxy) ExecuteCode(request ExecuteCodeRequest) ExecuteCodeResponse {
	fmt.Printf("代码沙箱请求消息:%v", request)
	executeCodeResponse := p.realSandbox.ExecuteCode(request)
	fmt.Printf("代码沙箱响应消息:%v", executeCodeResponse)
	return executeCodeResponse
}
