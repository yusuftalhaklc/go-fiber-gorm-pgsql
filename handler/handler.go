package handler

import (
	"go-fiber/database"
	"go-fiber/model"
	"go-fiber/util"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Create a user
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.DB
	user := new(model.User)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	user.Password = util.HashPassword(user.Password)
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

// Get All Users from db
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.DB
	var users []model.User

	db.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}

	token := c.Get("Authorization")
	if token == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "invalid Token"})
	}
	_, err := util.VerifyToken(token)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Unauthorized"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

// GetSingleUser from db
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.DB

	id := c.Params("id")
	var user model.User

	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

func UpdateUser(c *fiber.Ctx) error {
	type UpdateUser struct {
		Username string `json:"username"`
	}

	db := database.DB.DB
	var user model.User

	id := c.Params("id")

	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	var updateUserData UpdateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	user.Username = updateUserData.Username

	db.Save(&user)

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.DB
	var user model.User

	id := c.Params("id")

	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func LoginUser(c *fiber.Ctx) error {
	type LoginUser struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var login LoginUser
	db := database.DB.DB
	var user model.User

	err := c.BodyParser(&login)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	db.Find(&user, "username = ?", login.Username)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	if !(util.VerifyPassword(login.Password, user.Password)) {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	tokenString, err := util.CreateToken(&user)
	if err != nil {
		log.Println(tokenString)
		log.Println(err)
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Server Error"})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Login succesfuly", "token": tokenString})
}
