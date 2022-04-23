package book

import (
	_entities "book/entities"
)

type BookRepositoryInterface interface {
	GetAll() ([]_entities.Book, error)
	GetBookById(id int) (_entities.Book, int, error)
	CreateBook(book _entities.Book) error
	DeleteBook(id int) error
	UpdateBook(id int, book _entities.Book) error
}
