package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// @Summary Create a new question
// @Description Create a new question with the provided data
// @Tags Questions
// @Accept json
// @Produce json
// @Param question body models.Question true "Question Information"
// @Success 201 {object} models.Question
// @Failure 400 {error} error
// @Failure 500 {error} error
// @Router /api/v1/questions [post]
func CreateQuestion(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.Question{})
	var question models.Question
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
		"message":  "Question created successfully",
		"question": question,
	})
}

// @Summary Get all questions
// @Description Retrieve a list of all existing questions
// @Tags Questions
// @Accept json
// @Produce json
// @Success 200 {object} models.Question
// @Failure 500 {error} error
// @Router /api/v1/questions [get]
func GetQuestions(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.Question{})

	var questions []models.Question
	if err := global.Db.Find(&questions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":   "Questions retrieved successfully",
		"questions": questions,
	})
}

// @Summary Get a specific question by ID
// @Description Retrieve a question by its ID
// @Tags Questions
// @Accept json
// @Produce json
// @Param question_id path string true "Question ID"
// @Success 200 {object} models.Question
// @Failure 404 {error} error
// @Failure 500 {error} error
// @Router /api/v1/questions/{question_id} [get]
func GetQuestion(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.Question{})

	question_id := c.Params("question_id")

	var question models.Question

	if err := global.Db.Where("question_id = ?", question_id).First(&question).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":  "Question retrieved successfully",
		"question": question,
	})
}

// @Summary Update an existing question
// @Description Update a question based on its ID
// @Tags Questions
// @Accept json
// @Produce json
// @Param question_id path string true "Question ID"
// @Param question body models.Question true "Updated Question Information"
// @Success 200 {object} models.Question
// @Failure 400 {error} error
// @Failure 404 {error} error
// @Failure 500 {error} error
// @Router /api/v1/questions/{question_id} [put]
func UpdateQuestion(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.Question{})

	question_id := c.Params("question_id")

	var question models.Question
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
		"message":  "Question updated successfully",
		"question": question,
	})
}

// @Summary Delete a question by ID
// @Description Delete a specific question based on its ID
// @Tags Questions
// @Accept json
// @Produce json
// @Param question_id path string true "Question ID"
// @Success 200 {object} models.Question
// @Failure 500 {error} error
// @Router /api/v1/questions/{question_id} [delete]
func DeleteQuestion(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.Question{})
	question_id := c.Params("question_id")
	if err := global.Db.Where("question_id = ?", question_id).Delete(&models.Question{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Question deleted successfully",
	})
}