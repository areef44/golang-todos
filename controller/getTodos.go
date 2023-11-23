package controller

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
)

type TodoGetResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func GetTodosController(e *echo.Echo, db *sql.DB) {
	e.GET("/todos", func(ctx echo.Context) error {
		rows, err := db.Query("SELECT * FROM todos")

		if err != nil {
			return ctx.JSON(
				http.StatusInternalServerError,
				map[string]string{"error": err.Error()},
			)
		}

		var res []TodoGetResponse
		for rows.Next() {
			var id int64
			var title string
			var description string
			var done int
			err = rows.Scan(&id, &title, &description, &done)
			if err != nil {
				return ctx.JSON(
					http.StatusInternalServerError,
					map[string]string{"error": err.Error()},
				)
			}

			todo := TodoGetResponse{
				ID:          id,
				Title:       title,
				Description: description,
				Done:        done == 1,
			}

			res = append(res, todo)
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"message":  "Data Berhasil Didapatkan",
			"response": res,
		})
	})
}
