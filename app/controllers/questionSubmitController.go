package controllers

import (
	"OJ/app/models"
	"OJ/pkg/global"
	"OJ/pkg/utils"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

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
			"error": "无法解析",
		})
	}
	// 构建要发送的请求体
	requestBody := map[string]interface{}{
		"cmd": []map[string]interface{}{
			{
				"args": []string{"/usr/bin/g++", "a.cc", "-o", "a"},
				"env":  []string{"PATH=/usr/bin:/bin"},
				"files": []map[string]interface{}{
					{
						"content": "",
					},
					{
						"name": "stdout",
						"max":  10240,
					},
					{
						"name": "stderr",
						"max":  10240,
					},
				},
				"cpuLimit":    10000000000,
				"memoryLimit": 104857600,
				"procLimit":   50,
				"copyIn": map[string]map[string]string{
					"a.cc": {
						"content": question.Code,
					},
				},
				"copyOut":       []string{"stdout", "stderr"},
				"copyOutCached": []string{"a"},
			},
		},
	}

	// 将请求体编码为JSON
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 发送 POST 请求
	resp, err := http.Post("http://judge:5050/run", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Error sending POST request: %v", err)
		return c.Status(500).JSON(fiber.Map{
			"error": "请求发送错误",
		})
	}
	defer resp.Body.Close()

	type Response struct {
		Status     string `json:"status"`
		ExitStatus int    `json:"exitStatus"`
		Time       int64  `json:"time"`
		Memory     int64  `json:"memory"`
		RunTime    int64  `json:"runTime"`
		Files      struct {
			Stderr string `json:"stderr"`
			Stdout string `json:"stdout"`
		} `json:"files"`
		FileIds map[string]string `json:"fileIds"`
	}

	var response []Response

	// 读取原始响应体，帮助我们调试
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error reading response body",
		})
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Error decoding the response: " + err.Error(),
		})
	}

	// 将问题提交保存到数据库
	submission := models.QuestionSubmit{
		QuestionID: question.QuestionID,
		Code:       question.Code,
		Language:   question.Language,
		Status:     response[0].Status,
		ExitStatus: response[0].ExitStatus,
		Time:       response[0].Time,
		Memory:     response[0].Memory,
		RunTime:    response[0].RunTime,
		Stdout:     response[0].Files.Stdout,
		Stderr:     response[0].Files.Stderr,
	}

	// 将数据保存到数据库
	if err := global.Db.Create(&submission).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// 返回响应给前端
	// 返回成功消息
	return c.Status(201).JSON(fiber.Map{
		"message":  "QuestionSubmit created and run request sent successfully",
		"response": response,
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
