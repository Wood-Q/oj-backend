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

// @Summary Get a specific question submission by ID
// @Description Retrieve a question submission by its ID
// @Tags QuestionSubmits
// @Accept json
// @Produce json
// @Param question_id path string true "Question ID"
// @Success 200 {object} models.QuestionSubmit
// @Failure 404 {object} error
// @Router /api/v1/questionsubmit/{question_id} [get]
func GetQuestionSubmit(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.QuestionSubmit{})

	question_id := c.Params("question_id")

	var question models.QuestionSubmit

	if err := global.Db.Where("question_id = ?", question_id).First(&question).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	checkedUser := c.Locals("currentUser")
	if checkedUser.(*models.User).ID != question.UserID && checkedUser.(*models.User).UserRole != "admin" {
		return c.Status(401).JSON(fiber.Map{
			"error": "非提交用户，无法查看",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"message":  "QuestionSubmit retrieved successfully",
		"question": question,
	})
}

// @Summary Update an existing question submission
// @Description Update a question submission based on its ID
// @Tags QuestionSubmits
// @Accept json
// @Produce json
// @Param question_id path string true "Question ID"
// @Param questionSubmit body models.QuestionSubmit true "Updated QuestionSubmit Information"
// @Success 200 {object} models.QuestionSubmit
// @Failure 400 {object} error
// @Failure 404 {object} error
// @Failure 500 {object} error
// @Router /api/v1/questionsubmit/{question_id} [put]
func UpdateQuestionSubmit(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.QuestionSubmit{})

	question_id := c.Params("question_id")

	var question models.QuestionSubmit
	if err := global.Db.Where("question_id = ?", question_id).First(&question).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Update fields
	if err := c.BodyParser(&question); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := global.Db.Save(&question).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":  "QuestionSubmit updated successfully",
		"question": question,
	})
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
