package handler

import (
	s "GoGramTest/internal/state"
	"strconv"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func ShowCommandHandler(m *tg.NewMessage, bs *s.BotState) error {
	books, err := bs.Repo.GetAllBooks()
	if err != nil {
		return err
	}
	var result strings.Builder
	result.WriteString("<i><b>Сохраненные книги:</b></i> " + strconv.Itoa(len(books)) + "\n")
	for _, book := range books {
		result.WriteString(book.ToString())
	}

	_, err = m.Reply(result.String())
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}
	return nil
}
