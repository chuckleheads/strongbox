package main

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/labstack/echo"
	"net/http"
)

type Origin struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type NewOrigin struct {
	Name string `json:"name"`
}

func (o Origin) String() string {
	return fmt.Sprintf("Origin<%d %s>", o.Id, o.Name)
}

func main() {
	db := pg.Connect(&pg.Options{
		User:     "josh",
		Database: "strongbox",
	})
	defer db.Close()

	e := echo.New()

	// Create origin
	e.POST("/origins", func(c echo.Context) error {
		o := new(NewOrigin)
		if err := c.Bind(o); err != nil {
			return err
		}

		// This feels bad. There's probably a better pattern for this.
		origin := Origin{
			Name: o.Name,
		}

		err := db.Insert(&origin)

		if err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "")
		} else {
			return c.JSON(http.StatusCreated, origin)
		}
	})

	e.GET("/origins/:name", func(c echo.Context) error {
		name := c.Param("name")
		origin := new(Origin)

		err := db.Model(origin).
			Where("name = ?", name).
			Select()

		if err != nil {
			if err == pg.ErrNoRows {
				return c.String(http.StatusNotFound, "")
			} else {
				fmt.Println(err)
				return c.String(http.StatusInternalServerError, "")
			}
		} else {
			return c.JSON(http.StatusOK, origin)
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}
