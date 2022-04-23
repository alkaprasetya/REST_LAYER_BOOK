package book

import (
	"book/delivery/helper"
	_entities "book/entities"
	_bookUseCase "book/usecase/book"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	bookUseCase _bookUseCase.BookUseCaseInterface
}

func NewBookHandler(bookUseCase _bookUseCase.BookUseCaseInterface) *BookHandler {
	return &BookHandler{
		bookUseCase: bookUseCase,
	}
}

func (bh *BookHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		books, err := bh.bookUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all books", books))
	}
}

func (bh *BookHandler) GetByIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		books, rows, err := bh.bookUseCase.GetBookById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get books by id", books))
	}
}

func (bh *BookHandler) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var books _entities.Book
		c.Bind(&books)
		err := bh.bookUseCase.CreateBook(books)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to create book"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success Create book"))
	}
}

func (bh *BookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		err = bh.bookUseCase.DeleteBook(id)
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success delete book by id"))
	}
}

func (bh *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		var books _entities.Book
		c.Bind(&books)
		err = bh.bookUseCase.UpdateBook(id, books)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update book by id", books))
	}
}
