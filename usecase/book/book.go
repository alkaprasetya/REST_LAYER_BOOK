package book

import (
	_entities "book/entities"
	_bookRepository "book/repository/book"
)

type BookUseCase struct {
	BookRepository _bookRepository.BookRepositoryInterface
}

func NewBookUseCase(bookRepo _bookRepository.BookRepositoryInterface) BookUseCaseInterface {
	return &BookUseCase{
		BookRepository: bookRepo,
	}

}

func (buc *BookUseCase) GetAll() ([]_entities.Book, error) {
	books, err := buc.BookRepository.GetAll()
	return books, err
}

func (buc *BookUseCase) GetBookById(id int) (_entities.Book, int, error) {
	book, rows, err := buc.BookRepository.GetBookById(id)
	return book, rows, err
}

func (buc *BookUseCase) CreateBook(book _entities.Book) error {
	err := buc.BookRepository.CreateBook(book)
	return err
}

func (buc *BookUseCase) DeleteBook(id int) error {
	err := buc.BookRepository.DeleteBook(id)
	return err
}

func (buc *BookUseCase) UpdateBook(id int, book _entities.Book) error {
	err := buc.BookRepository.UpdateBook(id, book)
	return err
}
