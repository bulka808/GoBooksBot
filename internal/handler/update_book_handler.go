package handler

import (
	"GoGramTest/internal/model"
	s "GoGramTest/internal/state"
	"fmt"
	"log"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func UpdateBookHandler(m *tg.NewMessage, bs *s.BotState) error {
	defer bs.SetState(s.Idle)

	err := m.React("👍")
	if err != nil {
		log.Println(err)
		return err
	}

	lines := strings.Split(m.Text(), "\n")

	//делаем заглушки
	series, chapter := "---", "---"
	title := lines[0]
	author := strings.SplitN(lines[1], " ", 2)[1]
	if strings.Contains(m.Text(), "Серия:") {
		series = strings.SplitN(lines[2], " ", 2)[1]
	}
	if strings.Contains(m.Text(), "По:") {
		chapter = strings.SplitN(lines[4], " ", 2)[1]
	}

	newBook := model.Book{Title: title, Author: author, Series: series, Chapter: chapter}

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
	log.Println(book.ToString())
	// обновляем
	_, err = bs.Repo.Update(book, &newBook)
	// увеличиваем счетчик обновлений
	bs.AddUpdate(int(book.ID))
	if err != nil {
		log.Println(err)
		return err
	}

	time.Sleep(1 * time.Second)
	return nil
}
