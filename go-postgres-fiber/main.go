package main

import (
	"fmt"
	"go-postgres-fiber/models"
	"go-postgres-fiber/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// type Book struct {
// 	Author    string `json:"author"`
// 	Title     string `json:"title"`
// 	Publisher string `json:"publisher"`
// }

func (r *Repository) CreateBook(context *fiber.Ctx) error{
  book := models.Books{}

  err := context.BodyParser(&book) // This gets the book content like author, title, publisher from post request body and add it to book struct instance

  if err!=nil{
    context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
    return err;
  }
  err = r.DB.Create(&book).Error
  if err!=nil{
    context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message":"couldnot create book"})
    return err
  }
  context.Status(http.StatusOK).JSON(&fiber.Map{"message":"book has been added"})
  return nil
}

func (r *Repository) GetBooks(context *fiber.Ctx) error{
  bookModels := &[]models.Books{}

  err := r.DB.Find(bookModels).Error
  if err!= nil{
    context.Status(http.StatusBadRequest).JSON(
      &fiber.Map{"message":"could not get books"})
    return err
  }
  context.Status(http.StatusOK).JSON(&fiber.Map{
    "message": "Books fetched successfully",
    "data": bookModels,
  })
  return nil
}

func (r *Repository) DeleteBookById(context *fiber.Ctx) error{
  bookModel := models.Books{}
  id:= context.Params("id")
  if id == ""{
    context.Status(http.StatusInternalServerError).JSON(
      &fiber.Map{"message":"Couldnt delete book"},
      )
    return nil
  }

  err := r.DB.Delete(bookModel, id)

  if err.Error !=nil{
    context.Status(http.StatusBadRequest).JSON(&fiber.Map{
      "message":"not able to delete book"},
      )
    return err.Error
  }
  context.Status(http.StatusOK).JSON(&fiber.Map{
    "message": "Book Deleted successfully",
  })
  return nil
}

func (r *Repository) GetBookById(context *fiber.Ctx) error{
  bookModel := &models.Books{}
  id:= context.Params("id")
   if id == ""{
    context.Status(http.StatusInternalServerError).JSON(
      &fiber.Map{"message":"Couldnt delete book"},
      )
    return nil
  }

  fmt.Println("Id is: ",id)

  err := r.DB.Where("id = ?", id).First(bookModel).Error
  if err!=nil{
    context.Status(http.StatusBadRequest).JSON(
      &fiber.Map{"message":"couldnot get the book"},
      )
    return err
  }
  context.Status(http.StatusOK).JSON(
    &fiber.Map{"message": "Book fetched successfully", "data":bookModel},
    )
  return nil
 
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/create-books", r.CreateBook)
	api.Delete("/delete-book/:id", r.DeleteBookById)
	api.Get("/get-books/:id", r.GetBookById)
	api.Get("/books", r.GetBooks)
}


func setupApp() *fiber.App {
    // Initialize Fiber app
    app := fiber.New()

    // Database setup
    config := &storage.Config{
        Host: os.Getenv("DB_HOST"),
        Port: os.Getenv("DB_PORT"),
        Password: os.Getenv("DB_PASS"),
        User: os.Getenv("DB_USER"),
        SSLMode: os.Getenv("DB_SSLMode"),
        DBName: os.Getenv("DB_NAME"),
    }
    db, err := storage.NewConnection(config)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate models
    err = models.MigrateBooks(db)
    if err != nil {
        log.Fatal("Failed to migrate database models:", err)
    }

    // Setup routes
    r := Repository{DB: db}
    r.SetupRoutes(app)

    return app
}

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal(err)
    }

    app := setupApp()
    app.Listen(":8080")
}


// Improvements
// 1. Add database connection pooling, finetuneparameters based on server cpu,memory
// Wrapping app.listen() in goroutine doesnt make any faster. 
// app.listen() blocks the process in listens for requests. If you need to run any additional tasks after that, you can use goroutine to run it in background