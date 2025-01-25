package judge

import (
	"OJ/app/models"
	"OJ/pkg/enums"
)

type JudgeContext struct {
	JudgeInfo enums.JudgeInfo

	InputList  []string
	OutputList []string

	Question models.Question
	QuestionSubmit models.QuestionSubmit
}
