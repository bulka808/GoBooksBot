package handler

import (
	"GoGramTest/internal/model"
	s "GoGramTest/internal/state"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func AddBookHandler(m *tg.NewMessage, bs *s.BotState) error {
	defer bs.SetState(s.Idle)

	err := m.React("👍")
	if err != nil {
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

	// собираем книгу
	book := model.Book{
		BookId:  <-bs.AddId,
		Title:   title,
		Author:  author,
		Series:  series,
		Chapter: chapter,
		ChatId:  m.ChatID(),
	}

	_, err = bs.Repo.Create(&book)
	if err != nil {
		return err
	}

	return nil
}
