package main

import (
	"github.com/areef44/go-todos/controller"
	"github.com/areef44/go-todos/database"
	"github.com/labstack/echo"
)

func main() {

	db := database.ConnectDB()

	defer db.Close()

	err := db.Ping()

	if err != nil {
		panic(err)
	}

	e := echo.New()

	controller.PostTodosController(e, db)
	controller.GetTodosController(e, db)
	controller.DeleteTodosController(e, db)
	controller.UpdateTodosController(e, db)
	controller.CheckTodosController(e, db)

	e.Start(":8080")
}
