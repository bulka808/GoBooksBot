package handler

import (
	m "GoGramTest/internal/model"
	s "GoGramTest/internal/state"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func UpdateCommandHandler(m *tg.NewMessage, bs *s.BotState) error {
	bs.SetState(s.Update)

	err := m.React("👍")
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}

	books, err := bs.Repo.GetAllBooks()
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}

	var messages []*tg.NewMessage
	// -> UpdateBookHandler
	for _, book := range books {
		msg, err := m.Client.SendMessage(book.ChatId, getBookCommand(book))
		if err != nil {
			bs.SetState(s.Idle)
			return err
		}
		messages = append(messages, msg)
	}

	// удаляем сообщения
	go func() {
		// задержка на всякий случай
		time.Sleep(2 * time.Second)
		for _, msg := range messages {
			_, _ = msg.Delete()
		}
	}()

	//ждем пока словятся все книги
	time.Sleep(3 * time.Second)

	//собираем ответ
	var result strings.Builder
	result.WriteString("<i><b>Новое:</b></i> " + strconv.Itoa(len(bs.UpdateIDs)) + "\n")
	for id := range bs.UpdateIDs {
		book, err := bs.Repo.GetBookByID(id)
		if err != nil {
			log.Println(err)
		}
		result.WriteString(book.ToString())
	}

	result.WriteString("<i><b>Сохраненные книги:</b></i> " + strconv.Itoa(len(bs.UpdateIDs)) + "\n")
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

func getBookCommand(b *m.Book) string {
	return fmt.Sprintf("/download%d@botbybase_bot", b.BookId)
}
