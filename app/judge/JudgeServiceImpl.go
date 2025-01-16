package judge

import (
	"OJ/app/controllers"
	"OJ/app/models"
	"OJ/pkg/enums"
	"encoding/json"
	"errors"
)

const (
	ErrSubmitNotFound   = "提交不存在"
	ErrQuestionNotFound = "问题不存在"
	ErrJudgeInProgress  = "正在判题中"
	ErrUpdateFailed     = "更新失败"
	ErrParseFailed      = "解析失败"
)

type JudgeServiceImpl struct {
}

type JudgeCase struct {
	Input  string
	Output string
}

type JudgeConfig struct {
	TimeLimit   int
	MemoryLimit int
	StackLimit  int
}

func (j *JudgeServiceImpl) doJudge(questionSubmitId int64) (models.QuestionSubmit, error) {
	// 获取提交记录
	questionSubmit, err := controllers.GetQuestionSubmit(questionSubmitId)
	if err != nil {
		return questionSubmit, errors.New(ErrSubmitNotFound)
	}

	status := questionSubmit.Status
	if status != "WAITTING" {
		return questionSubmit, errors.New(ErrJudgeInProgress)
	}

	// 获取问题
	question, err := controllers.GetQuestion(questionSubmitId)
	if err != nil {
		return questionSubmit, errors.New(ErrQuestionNotFound)
	}

	// 更新为正在判题状态
	_, err = controllers.UpdateQuestionSubmit(questionSubmitId, "RUNNING")
	if err != nil {
		return questionSubmit, errors.New(ErrUpdateFailed)
	}

	// 调用代码沙箱执行代码
	factory := &enums.CodeSandboxFactory{}
	codeSandbox := factory.NewInstance("example")
	proxy := enums.NewCodeSandboxProxy(codeSandbox)

	var inputList []string
	err = json.Unmarshal([]byte(question.JudgeCase), &inputList)
	if err != nil {
		return questionSubmit, errors.New(ErrParseFailed)
	}

	// 执行代码请求
	executeCodeRequest := &enums.ExecuteCodeRequest{
		Code:      questionSubmit.Code,
		Language:  questionSubmit.Language,
		InputList: inputList,
	}
	executeCodeResponse := proxy.ExecuteCode(*executeCodeRequest)

	// 创建判断上下文
	judgeContext := &JudgeContext{
		InputList:  inputList,
		OutputList: executeCodeResponse.OutputList,
		JudgeInfo:  *executeCodeResponse.JudgeInfo,
	}

	// 执行判题逻辑
	judgeStrategy := &DefaultJudgeStrategy{}
	judgeStrategy.doJudge(*judgeContext)

	// 更新提交记录为成功
	_, err = controllers.UpdateQuestionSubmit(questionSubmitId, "SUCCESS")
	if err != nil {
		return questionSubmit, errors.New(ErrUpdateFailed)
	}

	// 获取最终的提交结果
	questionSubmitResult, err := controllers.GetQuestionSubmit(questionSubmitId)
	if err != nil {
		return questionSubmitResult, errors.New(ErrSubmitNotFound)
	}

	return questionSubmitResult, nil
}
