package controllers

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/hjuraev31/database"
	"github.com/hjuraev31/models"
	util "github.com/hjuraev31/utils"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}

func Register(ctx *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User
	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal("Couldn`t parse body")
	}

	if len(data["password"].(string)) <= 6 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Password length should be greater than 6 characters",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Invalid email address!",
		})
	}
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "User with this email already exists",
		})
	}
	user := models.User{
		FirstName: data["first_name"].(string),
		LastName:  data["last_name"].(string),
		Email:     strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}
	ctx.Status(200)
	return ctx.JSON(fiber.Map{
		"message": "User created successfully!",
		"user":    user,
	})
}

func Login(ctx *fiber.Ctx) error {
	var data map[string]string
	var user models.User

	if err := ctx.BodyParser(&data); err != nil {
		log.Println("Unable to parse body!")
	}

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		ctx.Status(404)
		return ctx.JSON(fiber.Map{
			"message": "User with this email doesn't exist",
		})
	}

	if err := user.ComparePassword(data["password"]); err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	token, err := util.GenerateJWT(strconv.Itoa(int(user.Id)))
	if err != nil {
		log.Println(err)
		ctx.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)

	return ctx.JSON(fiber.Map{
		"message": "Successfully logged in",
		"user":    user,
	})
}

type Claims struct {
	jwt.StandardClaims
}
