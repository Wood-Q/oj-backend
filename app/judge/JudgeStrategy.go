package judge

import "OJ/pkg/enums"

type JudgeStrategy interface {
	doJudge(judgeContext JudgeContext) enums.JudgeInfo
}