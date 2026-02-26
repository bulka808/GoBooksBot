package repository

import (
	"GoGramTest/internal/model"
	"context"

	"gorm.io/gorm"
)

type BookRepository struct {
	db  *gorm.DB
	ctx context.Context
}

func NewBookRepository(db *gorm.DB, ctx context.Context) *BookRepository {
	return &BookRepository{db: db, ctx: ctx}
}

func (repo *BookRepository) Create(book *model.Book) (*model.Book, error) {
	err := gorm.G[model.Book](repo.db).Create(repo.ctx, book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (repo *BookRepository) GetBookByID(id int) (*model.Book, error) {
	book, err := gorm.G[*model.Book](repo.db).Where("ID = ?", id).First(repo.ctx)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (repo *BookRepository) GetBookByBookID(bookID int) (*model.Book, error) {
	book, err := gorm.G[*model.Book](repo.db).Where("BookID = ?", bookID).First(repo.ctx)
	if err != nil {
		return nil, err
	}

	return book, nil
}
func (repo *BookRepository) GetAllByTitle(title string) ([]*model.Book, error) {
	books, err := gorm.G[*model.Book](repo.db).Where("Title LIKE ?", "%"+title+"%").Find(repo.ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
func (repo *BookRepository) GetAllByAuthor(author string) ([]*model.Book, error) {
	books, err := gorm.G[*model.Book](repo.db).Where("Author LIKE ?", "%"+author+"%").Find(repo.ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
func (repo *BookRepository) GetAllBySeries(series string) ([]*model.Book, error) {
	books, err := gorm.G[*model.Book](repo.db).Where("Series LIKE ?", "%"+series+"%").Find(repo.ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}
func (repo *BookRepository) GetAllBooks() ([]*model.Book, error) {
	books, err := gorm.G[*model.Book](repo.db).Find(repo.ctx)
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (repo *BookRepository) Update(book *model.Book, newBook *model.Book) (*model.Book, error) {
	_, err := gorm.G[model.Book](repo.db).Where("ID = ?", book.ID).Updates(repo.ctx, *newBook)
	if err != nil {
		return nil, err
	}
	return newBook, nil
}

func (repo *BookRepository) Delete(book *model.Book) (*model.Book, error) {
	_, err := gorm.G[*model.Book](repo.db).Where("ID = ?", book.ID).Delete(repo.ctx)
	if err != nil {
		return nil, err
	}

	return book, nil
}
