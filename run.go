package main

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"

	"golang-serialization-benchmark/src/model"
)

var (
	db   *gorm.DB
	User *model.User
)

func main() {
	var err error
	db, err = gorm.Open("mysql", "root:777777Mn!@tcp(192.168.0.107:3306)/golang?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}

	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))

	e.GET("/", hello)
	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	db.Where("id = ?", id).Find(User)
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	if err := c.Bind(User); err != nil {
		fmt.Println(err)
		return c.String(http.StatusBadRequest, "params is needed")
	}
	user := &model.User{Nickname:""}
	db.Create(user)
	return c.String(http.StatusOK, id)
}

func updateUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func deleteUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
