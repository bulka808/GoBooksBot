package handler

import (
	s "GoGramTest/internal/state"
	"log"
	"strconv"
	"strings"
	"time"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func AddCommandHandler(m *tg.NewMessage, bs *s.BotState) error {
	bs.SetState(s.Add)

	err := m.React("👍")
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}

	// /download123123@botbybase_bot
	// достаем BookId
	msg, err := m.GetReplyMessage()
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}
	log.Println(msg.Text())
	idStr := strings.Split(msg.Text()[9:], "@")[0]
	bookId, err := strconv.Atoi(idStr)
	if err != nil {
		bs.SetState(s.Idle)
		return err
	}
	log.Println(bookId)

	// присваиваем в состояние
	bs.AddId <- uint(bookId)
	// отправляем боту сообщение -> AddBookHandler
	log.Println(bookId)
	msg, err = m.Client.SendMessage(m.ChatID(), msg.Text())
	if err != nil {
		bs.SetState(s.Idle)
		<-bs.AddId
		return err
	}
	log.Println(msg.Text())

	// удаляем своё сообщение
	go func() {
		time.Sleep(2 * time.Second)
		_, _ = msg.Delete()
	}()

	return nil
}
