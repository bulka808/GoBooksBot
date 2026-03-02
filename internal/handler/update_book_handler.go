package handler

import (
	"GoGramTest/internal/model"
	s "GoGramTest/internal/state"
	"GoGramTest/internal/utils"
	"fmt"
	"log"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func UpdateBookHandler(m *tg.NewMessage, bs *s.BotState) error {
	err := m.React("👍")
	if err != nil {
		log.Println(err)
		return err
	}

	newBook := utils.ParseBookFromMessage(m)

	// достаем список всех книг
	books, err := bs.Repo.GetAllByTitle(newBook.Title)
	if err != nil {
		log.Println(err)
		return err
	}

	var book *model.Book

	// проходимся по списку книг, находим нужную
	for _, b := range books {
		if b.Title == newBook.Title && b.Author == newBook.Author {
			newBook.BookId = b.BookId
			newBook.ID = b.ID
			book = b
		}
	}

	// такого быть не должно, но лучше логировать
	if book == nil {
		log.Println("book not found!")
		return fmt.Errorf("book not found")
	}
	log.Println("parsed book: \n", book.ToString())
	// обновляем
	if newBook.Chapter != book.Chapter {
		_, err = bs.Repo.Update(book, &newBook)

		// добавляем id в список обновленных
		bs.AddUpdate(int(book.ID))
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
