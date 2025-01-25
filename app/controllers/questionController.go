package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	var question models.Question

	if err := c.BodyParser(&question); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "解析失败",
		})
	}

	utils.SetupDatabase(c, models.Question{})

	if err := global.Db.Create(&question).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err.Error(),
			"message": "创建失败",
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

// @Summary Get all questions with pagination
// @Description Retrieve a paginated list of all existing questions
// @Tags Questions
// @Accept json
// @Produce json
// @Param page query int false "Page number (default is 1)"
// @Param page_size query int false "Number of items per page (default is 10)"
// @Success 200 {object} models.Question
// @Failure 500 {error} error
// @Router /api/v1/questions/dividePage/questions [get]
func GetQuestionsByPage(c *fiber.Ctx) error {
	utils.SetupDatabase(c, models.Question{})

	// 获取分页参数，默认第一页和每页10条
	page := c.QueryInt("page", 1)
	pageSize := c.QueryInt("page_size", 10)

	log.Info(page, pageSize)

	// 获取过滤参数
	title := c.Query("title", "")

	tags := c.Query("tags", "")

	// 计算分页偏移量
	offset := (page - 1) * pageSize

	var questions []models.Question
	var totalCount int64

	// 创建查询构建器
	db := global.Db.Model(&models.Question{})

	// 如果有标题过滤
	if title != "" {
		db = db.Where("title LIKE ?", "%"+title+"%")
	}
	if tags != "" {
		// 将tags从字符串转换为数组
		tagList := strings.Split(tags, ",")
		for _, tag := range tagList {
			// 通过WHERE子句过滤tags字段
			db = db.Where("tags LIKE ?", "%"+tag+"%")
		}
	}

	// 获取总记录数
	if err := db.Count(&totalCount).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 获取分页数据
	if err := db.Limit(pageSize).Offset(offset).Find(&questions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message":     "Questions retrieved successfully",
		"questions":   questions,
		"total_count": totalCount, // 返回总记录数
		"page":        page,
		"page_size":   pageSize,
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
func GetQuestionByID(c *fiber.Ctx) error {
	var question models.Question
	question_id := c.Params("question_id")

	// 查询数据库
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

func GetQuestion(question_id int64) (models.Question, error) {
	var question models.Question

	// 查询数据库
	if err := global.Db.Where("question_id = ?", question_id).First(&question).Error; err != nil {
		return question, err
	}
	return question, nil
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
