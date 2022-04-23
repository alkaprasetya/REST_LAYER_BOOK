package main

import (
	"book/config"
	_userHandler "book/delivery/handler/user"
	_userRepository "book/repository/user"
	_userUseCase "book/usecase/user"

	_bookHandler "book/delivery/handler/book"
	_bookRepository "book/repository/book"
	_bookUseCase "book/usecase/book"

	"fmt"
	"log"

	_routes "book/delivery/routes"
	_utils "book/utils"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()
	db := _utils.InitDB(config)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	bookRepo := _bookRepository.NewBookRepository(db)
	bookUsecase := _bookUseCase.NewBookUseCase(bookRepo)
	bookHandler := _bookHandler.NewBookHandler(bookUsecase)

	e := echo.New()
	_routes.RegisterPath(e, userHandler, bookHandler)
	log.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}
