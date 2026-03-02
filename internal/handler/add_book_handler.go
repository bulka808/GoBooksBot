package handler

import (
	s "GoGramTest/internal/state"
	"GoGramTest/internal/utils"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func AddBookHandler(m *tg.NewMessage, bs *s.BotState) error {
	defer bs.SetState(s.Idle)

	err := m.React("👍")
	if err != nil {
		return err
	}

	// парсим книгу
	book := utils.ParseBookFromMessage(m)
	book.BookId = <-bs.AddId
	book.ChatId = m.ChatID()

	_, err = bs.Repo.Create(&book)
	if err != nil {
		return err
	}

	return nil
}
