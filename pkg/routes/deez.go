package routes

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleDeez(c echo.Context, db *sql.DB) error {
	var what string
	err := db.QueryRow("Select what from deez").Scan(&what)
	if err != nil {
		c.Logger().Errorf("Error querying database: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Database query error")
	}

	c.Logger().Infof("testing, testing.. 1,2: %v", what)
	return c.String(http.StatusOK, what)
}
