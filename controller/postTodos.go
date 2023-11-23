package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type CreateTodos struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func PostTodosController(e *echo.Echo, db *sql.DB) {
	e.POST("/todos", func(ctx echo.Context) error {
		var payload CreateTodos
		err := json.NewDecoder(ctx.Request().Body).Decode(&payload)
		if err != nil {
			return ctx.JSON(
				http.StatusBadRequest,
				map[string]string{"error": "Error decoding request body"},
			)
		}

		var id int64
		err = db.QueryRow(
			"INSERT INTO todos(title, description, done) VALUES ($1, $2, 0) RETURNING id",
			payload.Title,
			payload.Description,
		).Scan(&id)
		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": err.Error()},
			)
		}

		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": err.Error()},
			)
		}

		response := TodoResponse{
			ID:          id,
			Title:       payload.Title,
			Description: payload.Description,
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"message":  "Data berhasil ditambahkan",
			"response": response,
		})
	})
}
