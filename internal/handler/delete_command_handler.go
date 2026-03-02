package handler

import (
	s "GoGramTest/internal/state"
	"log"
	"strconv"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func DeleteCommandHandler(m *tg.NewMessage, bs *s.BotState) error {
	log.Println("delete command")
	err := m.React("👍")
	if err != nil {
		log.Println(err)
		return err
	}

	// /delete -> 7
	stringId := m.Text()[7:]
	id, err := strconv.Atoi(stringId)
	if err != nil {
		log.Println(err)
		return nil
	}

	book, err := bs.Repo.GetBookByID(id)
	if err != nil {
		log.Println(err)
		return nil
	}
	if book == nil {
		log.Println("book not found!")
		_, err = m.Reply("Book not found!")
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	}

	err = bs.Repo.DeleteById(id)
	if err != nil {
		log.Println(err)
		return nil
	}
	return nil
}
