package utils

import (
	"GoGramTest/internal/model"
	"fmt"
	"strconv"
	"strings"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func GetBookIdFromBotMessage(m *tg.NewMessage) (int, error) {
	idStr := strings.Split(m.Text()[9:], "@")[0]

	return strconv.Atoi(idStr)
}
func GetBookCommand(b *model.Book) string {
	return fmt.Sprintf("/download%d@botbybase_bot", b.BookId)
}

func ParseBookFromMessage(m *tg.NewMessage) model.Book {
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
	return model.Book{Title: title, Author: author, Series: series, Chapter: chapter}
}
