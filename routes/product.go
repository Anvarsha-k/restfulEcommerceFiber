package routes

import (
	"errors"

	"github.com/Anvarsha-k/restfulEcommerceFiber/database"
	"github.com/Anvarsha-k/restfulEcommerceFiber/models"
	"github.com/gofiber/fiber/v2"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func CreateProductResponse(productModel *models.Product)ProductSerializer {
	return ProductSerializer{ID: productModel.ID,Name: productModel.Name,SerialNumber: productModel.SerialNumber}
}

func CreateProduct(c *fiber.Ctx)error{
	var product models.Product

	if err :=c.BodyParser(&product); err!=nil{
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct:=CreateProductResponse(&product)
	return c.Status(200).JSON(responseProduct)
}
func GetProducts(c *fiber.Ctx)error{
	var products = []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts:=[]ProductSerializer{}

	for _,product:=range products{
		responseProduct:=CreateProductResponse(&product)
		responseProducts=append(responseProducts, responseProduct)

	}
	return c.Status(200).JSON(responseProducts)
}
//helper function
func FindProduct(id int,product *models.Product)error{
	database.Database.Db.Find(&product,"id=?",id)

	if product.ID == 0{
		return errors.New("product not exist")
	}
	return nil
}
func GetProduct(c *fiber.Ctx)error{
	id,err:=c.ParamsInt("id")

	var product models.Product
	if err != nil{
		return  c.Status(400).JSON("product is valid")
	}
	if err:=FindProduct(id,&product);err!= nil{
		return c.Status(400).JSON(err.Error())
	}
	responseProduct:=CreateProductResponse(&product)
	return c.Status(200).JSON(responseProduct)
}
func UpdateProduct(c *fiber.Ctx)error{
	id,err := c.ParamsInt("id")
	var product models.Product
	 if err!= nil{
		return c.Status(400).JSON("ensure product is valid")
	 }
	 if err:=FindProduct(id,&product);err!=nil{
		return c.Status(400).JSON(err.Error())
	 }

	 type UpdateProducts struct{
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	 }

	 var UpdateData UpdateProducts

	 if err:=c.BodyParser(&UpdateData); err != nil{
		return c.Status(400).JSON(err.Error())
	 }

	 product.Name = UpdateData.Name
	 product.SerialNumber = UpdateData.SerialNumber

	 database.Database.Db.Save(&product)
	 responseProduct := CreateProductResponse(&product)
	 return c.Status(200).JSON(responseProduct)
	 
}

func DeleteProduct(c *fiber.Ctx)error{
	var product models.Product

	id,err:=c.ParamsInt("id")

	if err!= nil{
		return c.Status(400).JSON("ensure product is valid")
	}
	if err:=FindProduct(id,&product);err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	if err:=database.Database.Db.Delete(&product).Error;err!=nil{
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("Product deleted Successfully")
	
}