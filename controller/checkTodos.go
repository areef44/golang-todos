package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CheckTodos struct {
	Done bool `json:"done"`
}

type TodoCheckResponse struct {
	ID   int64 `json:"id"`
	Done bool  `json:"done"`
}

func CheckTodosController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id/check", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var payload CheckTodos
		if err := json.NewDecoder(ctx.Request().Body).Decode(&payload); err != nil {
			return ctx.JSON(
				http.StatusBadRequest,
				map[string]string{"error": "Error decoding request body"},
			)
		}

		var doneInt int

		if payload.Done {
			doneInt = 1
		}

		_, err := db.Exec(
			"UPDATE todos SET done=$1 WHERE id=$2",
			doneInt,
			id,
		)
		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": err.Error()},
			)
		}

		updatedID, err := strconv.Atoi(id)
		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": "Error converting ID to integer"},
			)
		}

		response := TodoCheckResponse{
			ID:   int64(updatedID),
			Done: doneInt == 1,
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"message":  "Data berhasil diupdate",
			"response": response,
		})
	})
}
