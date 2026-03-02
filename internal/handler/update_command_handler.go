package handler

import (
	s "GoGramTest/internal/state"
	"GoGramTest/internal/utils"
	"log"
	"strconv"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func UpdateCommandHandler(m *tg.NewMessage, bs *s.BotState) error {
	defer bs.SetState(s.Idle)
	bs.SetState(s.Update)
	log.Println("start update")

	err := m.React("👍")
	if err != nil {
		return err
	}

	// создаем список всех книг
	books, err := bs.Repo.GetAll()
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}

	log.Println("wait books")
	var messages []*tg.NewMessage
	// отправляем сообщения, их ловит -> UpdateBookHandler
	for _, book := range books {
		msg, err := m.Client.SendMessage(book.ChatId, utils.GetBookCommand(book))
		if err != nil {
			log.Println(err)
			return err
		}
		messages = append(messages, msg)
	}

	// удаляем сообщения
	go func() {
		// задержка на всякий случай
		time.Sleep(3 * time.Second)
		for _, msg := range messages {
			_, _ = msg.Delete()
		}
	}()

	//ждем пока словятся все книги
	time.Sleep(3 * time.Second)
	log.Println("ready")

	//собираем ответ
	log.Println("===building response===")
	log.Println("updated books:")
	var result strings.Builder
	result.WriteString("<i><b>Новое:</b></i> " + strconv.Itoa(len(bs.UpdateIDs)) + "\n")

	for range len(bs.UpdateIDs) {
		book, err := bs.Repo.GetBookByID(<-bs.UpdateIDs)
		if err != nil {
			log.Println(err)
		}

		log.Println(book.ID, book.Title)

		result.WriteString(book.Format())
	}

	log.Println("stored books: ")
	result.WriteString("<i><b>Сохраненные книги:</b></i> " + strconv.Itoa(len(books)) + "\n")
	for _, book := range books {
		log.Println(book.ID, book.Title)
		result.WriteString(book.Format())
	}

	_, err = m.Reply(result.String())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
