package main

import (
	"database/sql"
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qsliu2017/good/web"
)

type server struct {
	*echo.Echo
	*sql.DB
}

//go:embed schema.sql
var schema string

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		panic(err)
	}

	if _, err := db.Exec(schema); err != nil {
		panic(err)
	}

	s := &server{
		Echo: e,
		DB:   db,
	}

	s.StaticFS("/", web.StaticFs())

	s.GET("/api/user/:id", s.getUser)
	s.POST("/api/user", s.createUser)

	if err := s.Start(":8901"); err != nil {
		panic(err)
	}
}

type user struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (s *server) getUser(c echo.Context) error {
	var user user
	if err := s.QueryRowContext(c.Request().Context(),
		`SELECT
			id,
			username
		FROM user
		WHERE id = $1`,
		c.Param("id"),
	).Scan(
		&user.ID,
		&user.Username,
	); err == sql.ErrNoRows {
		return c.NoContent(http.StatusNotFound)
	} else if err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

type userCreate struct {
	Username string `json:"username"`
}

func (s *server) createUser(c echo.Context) error {
	var create userCreate
	if err := c.Bind(&create); err != nil {
		c.Logger().Error(err)
		return err
	}

	var user user
	if err := s.QueryRowContext(c.Request().Context(),
		`INSERT INTO user (
			username
		) VALUES (
			$1
		) RETURNING
			id,
			username`,
		create.Username,
	).Scan(
		&user.ID,
		&user.Username,
	); err != nil {
		c.Logger().Error(err)
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
