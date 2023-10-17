package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbname = "hrms"
const mongoURI = "mongodb://localhost:27017" + dbname

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {
	mongo.NewClient(options.Client().ApplyURI(mongoURI))
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {

	})
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")

}
