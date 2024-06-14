package controller

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"todo-app/config"
	"todo-app/helper"
	"todo-app/model/entity"
	"todo-app/model/request"
)

func GetAll(c *fiber.Ctx) error {
	var todo []entity.Todo

	err := config.DB.Find(&todo).Error
	if err != nil {
		return helper.Response(c, http.StatusInternalServerError, "Failed get data", nil)
	}

	return helper.Response(c, http.StatusOK, "Get all data", todo)
}

func Create(c *fiber.Ctx) error {
	var requestTodo request.RequestCreateTodo

	errParser := c.BodyParser(&requestTodo)
	if errParser != nil {
		return helper.Response(c, http.StatusBadRequest, "Failed body request", nil)
	}

	todo := entity.Todo{
		Title:       requestTodo.Title,
		Description: requestTodo.Description,
	}

	if err := config.DB.Create(&todo).Error; err != nil {
		return helper.Response(c, http.StatusInternalServerError, "Failed create data", nil)
	}

	return helper.Response(c, http.StatusOK, "Create todo successfully", requestTodo)
}

func GetById(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo entity.Todo

	err := config.DB.First(&todo, "id=?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.Response(c, http.StatusNotFound, "Data not found", nil)
	}

	return helper.Response(c, http.StatusOK, "Get data by id", todo)
}

func Update(c *fiber.Ctx) error {
	var requestTodo request.RequestCreateTodo
	errParser := c.BodyParser(&requestTodo)
	if errParser != nil {
		return helper.Response(c, http.StatusBadRequest, "Failed body request", nil)
	}

	id := c.Params("id")
	var existingTodo entity.Todo
	err := config.DB.First(&existingTodo, "id=?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.Response(c, http.StatusNotFound, "Data not found", nil)
	}

	existingTodo.Title = requestTodo.Title
	existingTodo.Description = requestTodo.Description

	if err := config.DB.Save(&existingTodo).Error; err != nil {
		return helper.Response(c, http.StatusInternalServerError, "Failed create data", nil)
	}

	return helper.Response(c, http.StatusOK, "Update todo successfully", existingTodo)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo entity.Todo
	err := config.DB.First(&todo, "id=?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return helper.Response(c, http.StatusNotFound, "Data not found", nil)
	}

	if err := config.DB.Where("id=?", id).Delete(&todo).Error; err != nil {
		return helper.Response(c, http.StatusInternalServerError, "Failed Delete data", nil)
	}

	return helper.Response(c, http.StatusOK, "Delete todo successfully", nil)
}
