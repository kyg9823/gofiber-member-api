package handlers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/kyg9823/gofiber-member-api/config"
	"github.com/kyg9823/gofiber-member-api/database"
	"github.com/kyg9823/gofiber-member-api/model"

	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {

	loginRequest := new(model.User)
	if err := ctx.BodyParser(loginRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Can't parse user information")
	}

	userInfo, err := getUserInfoByUsername(loginRequest.Username)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Can't parse user information")
	}

	log.Printf("%v", userInfo)
	isValidPassword, _ := comparePassword(userInfo.Password, loginRequest.Password)

	if !isValidPassword {
		return fiber.NewError(fiber.StatusUnauthorized, "Failed to login")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userInfo.Username
	claims["name"] = userInfo.Name
	claims["expired_date"] = time.Now().Add(time.Hour * 72).Unix()

	resultToken, err := token.SignedString([]byte(config.GetConfig("ACCESS_SECRET")))
	if err != nil {
		log.Printf("%v", err)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status":       200,
		"access_token": resultToken,
		"token_type":   "bearer",
	})
}

func hashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func comparePassword(hashedPassword, plainPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return false, err
		}
		return false, err
	}
	return true, nil
}

func getUserInfoByUsername(username string) (*model.User, error) {
	userInfo := new(model.User)
	result := database.DBConn.Table("users").Where("username == ?", username).First(userInfo)
	return userInfo, result.Error
}
