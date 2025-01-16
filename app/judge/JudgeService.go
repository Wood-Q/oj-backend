package judge

import "OJ/pkg/enums"

type JudgeService interface {
	doJudege(questionSubmitId int64) enums.ExecuteCodeResponse
}
