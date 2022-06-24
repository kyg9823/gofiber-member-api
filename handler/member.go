package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kyg9823/gofiber-member-api/database"
	"github.com/kyg9823/gofiber-member-api/model"
)

func GetMembers(ctx *fiber.Ctx) error {
	var members []model.Member
	database.DBConn.Find(&members)
	return ctx.JSON(members)
}

func GetMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")

	var member model.Member
	result := database.DBConn.First(&member, memberId)
	log.Printf("%v", result)
	if result.RowsAffected == 0 {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"status":  404,
			"message": "Not Found",
		})
	}

	return ctx.JSON(member)
}

func NewMember(ctx *fiber.Ctx) error {
	member := new(model.Member)
	if err := ctx.BodyParser(member); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	database.DBConn.Create(&member)
	return ctx.JSON(member)
}

func PutMember(ctx *fiber.Ctx) error {
	member := new(model.Member)
	if err := ctx.BodyParser(member); err != nil {
		return ctx.Status(503).SendString(err.Error())
	}
	database.DBConn.Create(&member)
	return ctx.JSON(fiber.Map{
		"status": 200,
		"result": "OK",
	})
}

func DeleteMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")
	var member model.Member
	result := database.DBConn.First(&member, memberId)
	log.Printf("%v", result)
	if result.RowsAffected == 0 {
		ctx.Status(204)
		return ctx.JSON(fiber.Map{
			"status":  204,
			"message": "No Content",
		})
	}

	result = database.DBConn.Delete(&member)
	if result.Error != nil {
		log.Printf("Error")
	}

	return ctx.JSON(fiber.Map{
		"status": 200,
		"result": "OK",
	})
}
