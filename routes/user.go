package routes

import (
	"errors"

	"github.com/Anvarsha-k/restfulEcommerceFiber/database"
	"github.com/Anvarsha-k/restfulEcommerceFiber/models"
	"github.com/gofiber/fiber/v2"
)

type UserSerializer struct {

	///this is not a model its just a serializer
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

func CreateResponseUser(userModel models.User) UserSerializer {

	//🔸 CreateResponseUser(user)
	// This is (most likely) a function

	// Takes a full models.User

	// Returns a filtered and formatted version of the user for response

	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, Lastname: userModel.Lastname}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	} // in here actually the bodyparser is a method that read the request and automatically parses into given go struct that is user

	database.Database.Db.Create(&user)       // inputting value to the database
	responseUser := CreateResponseUser(user) //creating the response of data maybe

	return c.Status(201).JSON(responseUser)

}	
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{} //This line creates an empty list (slice) to hold the cleaned-up users that you'll return to the frontend.

	for _, user := range users {
		responseUser := CreateResponseUser(user) //taking only needed data instead of givnig all sensitive data like passwords etc

		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

// helper function for individual user fetching
func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user not exist")
	}
	return nil
}
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User

	if err != nil {
		return c.Status(400).JSON("ensure user is valid")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User
	if err != nil {
		return c.Status(400).JSON("ensure user is valid")
	}

	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		Lastname  string `json:"last_name"`
	}

	var UpdateData UpdateUser

	if err := c.BodyParser(&UpdateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.FirstName = UpdateData.FirstName
	user.Lastname = UpdateData.Lastname

	database.Database.Db.Save(&user)

	responseUSer := CreateResponseUser(user)
	return c.Status(200).JSON(responseUSer)
}
func DeleteUser(c *fiber.Ctx)error{
	id,err:=c.ParamsInt("id")
	var user models.User

	if err!= nil {
		return c.Status(400).JSON("ensure a valid user")
	}	
	if err:=FindUser(id,&user);err!=nil{
		return c.Status(400).JSON(err.Error())
	}
	if err:=database.Database.Db.Delete(&user).Error;err!=nil{
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("user deleted successfully")
}