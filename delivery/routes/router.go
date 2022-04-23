package routes

import (
	_bookHandler "book/delivery/handler/book"
	_userHandler "book/delivery/handler/user"

	"github.com/labstack/echo/v4"
)

func RegisterPath(e *echo.Echo, uh *_userHandler.UserHandler, bh *_bookHandler.BookHandler) {
	e.GET("/users", uh.GetAllHandler())
	e.GET("/users/:id", uh.GetByIdHandler())
	e.POST("/users", uh.CreateUser())
	e.PUT("/users/:id", uh.UpdateUser())
	e.DELETE("/users/:id", uh.DeleteUser())

	e.GET("/books", bh.GetAllHandler())
	e.GET("/books/:id", bh.GetByIdHandler())
	e.POST("/books", bh.CreateBook())
	e.PUT("/books/:id", bh.UpdateBook())
	e.DELETE("/books/:id", bh.DeleteBook())
}
