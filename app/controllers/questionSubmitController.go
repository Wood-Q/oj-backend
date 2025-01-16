package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary Create a new question submit
// @Description Create a new question submission with the provided data
// @Tags QuestionSubmits
// @Accept json
// @Produce json
// @Param questionSubmit body models.QuestionSubmit true "Question Submit Information"
// @Success 201 {object} models.QuestionSubmit
// @Failure 400 {object} error
// @Failure 500 {object} error
// @Router /api/v1/questionsubmit [post]
func CreateQuestionSubmit(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.QuestionSubmit{})
	var question models.QuestionSubmit
	if err := c.BodyParser(&question); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if err := global.Db.Create(&question).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(201).JSON(fiber.Map{
		"message":  "QuestionSubmit created successfully",
		"question": question,
	})
}

// @Summary Get all question submissions
// @Description Retrieve a list of all existing question submissions
// @Tags QuestionSubmits
// @Accept json
// @Produce json
// @Success 200 {object} models.QuestionSubmit
// @Failure 500 {object} error
// @Router /api/v1/questionsubmit [get]
func GetQuestionSubmits(c *fiber.Ctx) error {

	utils.SetupDatabase(c, models.QuestionSubmit{})
	var questions []models.QuestionSubmit
	if err := global.Db.Find(&questions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":   "QuestionSubmits retrieved successfully",
		"questions": questions,
	})
}

func GetQuestionSubmit(question_id int64) (models.QuestionSubmit, error) {
	var question models.QuestionSubmit

	// 查询数据库
	if err := global.Db.Where("question_id = ?", question_id).First(&question).Error; err != nil {
		return question, err
	}
	return question, nil
}

func UpdateQuestionSubmit(question_id int64, status string) (models.QuestionSubmit, error) {

	var question models.QuestionSubmit
	if err := global.Db.Where("question_id = ?", question_id).First(&question).Error; err != nil {
		return question, err
	}

	question.Status = status

	if err := global.Db.Save(&question).Error; err != nil {
		return question, err
	}

	return question, nil
}

// @Summary Delete a question submission by ID
// @Description Delete a specific question submission based on its ID
// @Tags QuestionSubmits
// @Accept json
// @Produce json
// @Param question_id path string true "Question ID"
// @Success 200 {object} models.QuestionSubmit
// @Failure 500 {error} error
// @Router /api/v1/questionsubmit/{question_id} [delete]
func DeleteQuestionSubmit(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.QuestionSubmit{})
	question_id := c.Params("question_id")
	if err := global.Db.Where("question_id = ?", question_id).Delete(&models.QuestionSubmit{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "QuestionSubmit deleted successfully",
	})
}
