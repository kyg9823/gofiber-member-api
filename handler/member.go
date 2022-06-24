package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kyg9823/gofiber-member-api/database"
	"github.com/kyg9823/gofiber-member-api/model"
)

func GetMembers(ctx *fiber.Ctx) error {
	var members []model.MemberResponse
	database.DBConn.Table("members").Select("id", "name", "email").Find(&members)
	return ctx.Status(fiber.StatusOK).JSON(members)
}

func GetMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")

	var member model.MemberDetailResponse
	result := database.DBConn.
		Table("members").
		Select("members.id", "members.name", "members.email", "group_concat(favorites.item) AS favorites").
		Joins("left join favorites on members.id = favorites.id").Where("members.id = ?", memberId).
		Group("members.id, members.name, members.email").
		Find(&member)

	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": fiber.StatusNotFound,
			"result": "Not Found",
		})
	}

	if result.Error != nil {
		return result.Error
	}

	return ctx.Status(fiber.StatusOK).JSON(member)
}

func NewMember(ctx *fiber.Ctx) error {
	memberId, _ := ctx.ParamsInt("id")
	memberRequest := new(model.MemberRequest)
	if err := ctx.BodyParser(memberRequest); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can't parse body data.")
	}

	member := &model.Member{
		Id: int32(memberId),
	}
	member.ConvertFromRequest(memberRequest)
	result := database.DBConn.Create(&member)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(memberRequest)
}

func PutMember(ctx *fiber.Ctx) error {
	member := new(model.Member)
	if err := ctx.BodyParser(member); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can't parse body data.")
	}
	result := database.DBConn.Save(&member)

	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status": fiber.StatusNoContent,
			"result": "No Content",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": "OK",
	})
}

func DeleteMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")
	var member model.Member
	result := database.DBConn.First(&member, memberId)
	if result.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"status": fiber.StatusNoContent,
			"result": "No Content",
		})
	}

	result = database.DBConn.Delete(&member)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": "OK",
	})
}
