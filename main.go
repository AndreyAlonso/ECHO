package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover()) //permite registrar cualquier middlerware

	e.GET("/", saludar)
	e.GET("/dividir", dividir)

	// Grupo de rutas
	persons := e.Group("/personas")

	// Se asigna el middleware a utilizar para el grupo de rutas
	persons.Use(middlewareLogPersonas)

	persons.POST("", crear)
	persons.DELETE("", borrar)
	persons.GET("/:id", consultar) //localhost:8080/personas/2
	persons.PUT("/:id", actualizar)
	e.Start(":8080")
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "Hola Mundo"})
}

func dividir(c echo.Context) error {
	d := c.QueryParam("d") //QueryParam obtiene el parámetro de la URL

	f, _ := strconv.Atoi(d)
	if f == 0 {
		return c.String(http.StatusBadRequest, "\nEl valor no puede ser cero\n")
	}
	r := 3000 / f
	return c.String(http.StatusOK, strconv.Itoa(r))
}

func crear(c echo.Context) error {
	return c.String(http.StatusOK, "creado")
}

func actualizar(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "actualizado "+id)
}

func borrar(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "borrado "+id)
}

func consultar(c echo.Context) error {
	id := c.Param("id")

	return c.String(http.StatusOK, "consultado "+id)
}

func middlewareLogPersonas(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("Petición hecha a /personas")
		return f(c)
	}
}
