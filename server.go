package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func main() {
	// DB init
	var err error
	db, err = gorm.Open(sqlite.Open("quiz.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Question{})

	// Echo init∏
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	//e.Use(middleware.Recover())
	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	})
	e.POST("/quiz/", postAddQuestion)
	e.GET("/quiz/quess_movie", getQuessMovie)
	e.Logger.Fatal(e.Start(":8080"))
}

func postAddQuestion(c echo.Context) (err error) {
	question := new(Question)
	if err = c.Bind(question); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if result := db.Create(&question); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "DB ERR")
	}

	return c.JSON(http.StatusOK, question)
}

func getQuessMovie(c echo.Context) error {
	var question Question

	if result := db.Order("RANDOM()").Limit(1).First(&question); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, "DB ERR")
	}

	return c.JSON(http.StatusOK, question)
}
