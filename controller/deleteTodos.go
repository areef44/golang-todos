package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

func DeleteTodosController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		_, err := db.Exec(
			"DELETE FROM todos WHERE id = $1",
			id,
		)

		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": err.Error()},
			)
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Data berhasil dihapus",
		})
	})
}
