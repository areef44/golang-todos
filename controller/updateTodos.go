package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UpdateTodos struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TodoUpdateResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func UpdateTodosController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var payload UpdateTodos
		err := json.NewDecoder(ctx.Request().Body).Decode(&payload)
		if err != nil {
			return ctx.JSON(
				http.StatusBadRequest,
				map[string]string{"error": "Error decoding request body"},
			)
		}

		err = db.QueryRow(
			"UPDATE todos SET title=$1, description=$2 WHERE id=$3 returning id",
			payload.Title,
			payload.Description,
			id,
		).Scan(&id)

		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": err.Error()},
			)
		}

		updatedID, err := strconv.Atoi(id)

		response := TodoUpdateResponse{
			ID:          int64(updatedID),
			Title:       payload.Title,
			Description: payload.Description,
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"message":  "Data berhasil diupdate",
			"response": response,
		})
	})
}
