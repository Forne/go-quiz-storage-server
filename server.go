package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	// DB init
	var err error
	configDatabaseDriver, existsDriver := os.LookupEnv("DATABASE_DRIVER")
	configDatabaseDSN, existsDSN := os.LookupEnv("DATABASE_DSN")
	if !existsDriver || !existsDSN {
		panic("Config: DSN is not set")
	}

	if configDatabaseDriver == "sqlite" {
		db, err = gorm.Open(sqlite.Open(configDatabaseDSN), &gorm.Config{})
	} else if configDatabaseDriver == "mysql" {
		db, err = gorm.Open(mysql.Open(configDatabaseDSN), &gorm.Config{})
	} else {
		panic("Config: Database driver is not set")
	}

	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate
	configDatabaseMigration, exists := os.LookupEnv("DATABASE_MIGRATE")
	if exists && configDatabaseMigration == "true" {
		log.Printf("Config: Auto-migrate is enabled")
		db.AutoMigrate(&Locale{})
		db.AutoMigrate(&Category{})
		db.AutoMigrate(&Question{})
		db.AutoMigrate(&Answer{})
	}

	// Echo init
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.Static("/static", "static")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	//e.POST("/questions", postQuestion) // TODO
	e.GET("/questions", getQuestions)
	e.GET("/categories", getCategories)
	e.GET("/locales", getLocales)
	e.Logger.Fatal(e.Start(":8080"))
}

func postQuestion(c echo.Context) (err error) {
	question := new(Question)
	if err = c.Bind(question); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if result := db.Create(&question); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "DB ERR")
	}

	return c.JSON(http.StatusOK, question)
}

func getQuestions(c echo.Context) error {
	var question []Question

	query := db.Preload("Answers").Order("RANDOM()")

	// Category
	categoryId, err := strconv.Atoi(c.QueryParam("category_id"))
	if err == nil && categoryId != 0 {
		query = query.Where("category_id = ?", categoryId)
	}

	// Question type
	typeId, err := strconv.Atoi(c.QueryParam("type_id"))
	if err == nil && typeId != 0 {
		query = query.Where("type_id = ?", typeId)
	}

	// Locale
	localeCode := c.QueryParam("locale_code")
	if localeCode != "" {
		query = query.Where("locale_code = ?", localeCode)
	}

	// Limit
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err == nil && limit != 0 {
		query = query.Limit(limit)
	} else {
		query = query.Limit(10)
	}

	if result := query.Debug().Find(&question); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "DB ERR")
	}

	return c.JSON(http.StatusOK, question)
}

func getCategories(c echo.Context) error {
	var categories []Category

	if result := db.Find(&categories); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "DB ERR")
	}

	return c.JSON(http.StatusOK, categories)
}

func getLocales(c echo.Context) error {
	var locales []Locale

	if result := db.Find(&locales); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "DB ERR")
	}

	return c.JSON(http.StatusOK, locales)
}
