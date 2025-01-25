package judge

import "OJ/pkg/enums"

type JudgeManager struct{}

func (jm JudgeManager) DoJudge(judgeContext JudgeContext) enums.JudgeInfo {
	questionSubmit := judgeContext.QuestionSubmit
	var judgeStrategy JudgeStrategy = &DefaultJudgeStrategy{}
	language := questionSubmit.Language
	if language == "java" {
		// judgeStrategy = &JavaLanguageJudgeStrategy{}
	}
	return judgeStrategy.doJudge(judgeContext)
}
