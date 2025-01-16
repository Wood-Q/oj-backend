package judge

import (
	"OJ/pkg/enums"
	"encoding/json"
)

type DefaultJudgeStrategy struct {
}

func (j *DefaultJudgeStrategy) doJudge(judgeContext JudgeContext) enums.JudgeInfo {

	question := judgeContext.Question
	inputList := judgeContext.InputList
	outputList := judgeContext.OutputList
	judegeInfo := judgeContext.JudgeInfo
	memory := judegeInfo.Memory
	time := judegeInfo.Time
	judgeInfoResponse := new(enums.JudgeInfo)
	judgeInfoResponse.Message = "ACCEPTED"
	judegeInfo.Time = time
	judegeInfo.Memory = memory
	if len(outputList) != len(inputList) {
		judegeInfo.Message = "WRONG ANSWER"

		judgeInfoResponse.Message = judegeInfo.Message
		return *judgeInfoResponse
	}
	for i := 0; i < len(outputList); i++ {
		if outputList[i] != inputList[i] {
			judegeInfo.Message = "WRONG ANSWER"
			judgeInfoResponse.Message = judegeInfo.Message
			return *judgeInfoResponse
		}
	}
	//判断题目限制
	var judgeConfig JudgeConfig
	err := json.Unmarshal([]byte(question.JudgeConfig), &judgeConfig)
	if err != nil {
		return *judgeInfoResponse
	}
	memoryLimit := judgeConfig.MemoryLimit
	timeLimit := judgeConfig.TimeLimit

	if memory > memoryLimit {
		judegeInfo.Message = "MEMORY LIMIT EXCEEDED"
		judgeInfoResponse.Message = judegeInfo.Message
		return *judgeInfoResponse
	}
	if time > timeLimit {
		judegeInfo.Message = "TIME LIMIT EXCEEDED"
		judgeInfoResponse.Message = judegeInfo.Message
		return *judgeInfoResponse
	}

	return *judgeInfoResponse
}
