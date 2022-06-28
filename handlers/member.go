package handlers

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

	var member model.MemberDetail
	result := database.DBConn.
		Table("members").
		Select("members.id", "members.name", "members.email", "group_concat(favorites.item) as favorites").
		Joins("left join favorites on members.id = favorites.id").
		Where("members.id = ?", memberId).
		Group("members.id, members.name, members.email").
		Find(&member)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNotFound, "Not Found")
	}

	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	memberDetailResponse := new(model.MemberDetailResponse)
	memberDetailResponse.ConvertToResponse(&member)

	return ctx.Status(fiber.StatusOK).JSON(memberDetailResponse)
}

func NewMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")
	memberRequest := new(model.MemberRequest)
	if err := ctx.BodyParser(memberRequest); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can't parse body data.")
	}

	member := &model.Member{
		Id: memberId,
	}
	member.ConvertFromRequest(memberRequest)
	result := database.DBConn.Create(&member)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}
	return ctx.Status(fiber.StatusOK).JSON(memberRequest)
}

func PutMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")
	memberRequest := new(model.MemberRequest)
	if err := ctx.BodyParser(memberRequest); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can't parse body data.")
	}

	member := &model.Member{
		Id: memberId,
	}

	result := database.DBConn.Where("id = ?", memberId).Delete(&model.Favorite{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	member.ConvertFromRequest(memberRequest)
	result = database.DBConn.Save(&member)

	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNoContent, "Not Content")
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": "OK",
	})
}

func DeleteMember(ctx *fiber.Ctx) error {
	memberId := ctx.Params("id")
	var member model.Member
	result := database.DBConn.Where("id = ?", memberId).First(&member)
	if result.RowsAffected == 0 {
		return fiber.NewError(fiber.StatusNoContent, "Not Content")
	}

	result = database.DBConn.Where("id = ?", memberId).Delete(&member)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	result = database.DBConn.Where("id = ?", memberId).Delete(&model.Favorite{})
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, result.Error.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": fiber.StatusOK,
		"result": "OK",
	})
}
