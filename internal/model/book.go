package model

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)
import _ "gorm.io/driver/sqlite"

type Book struct {
	gorm.Model
	BookId  uint   `json:"book_id"`
	Title   string `gorm:"type:varchar(255);" json:"title"`
	Author  string `gorm:"type:varchar(255);" json:"author"`
	Series  string `gorm:"type:varchar(255);" json:"series"`
	Chapter string `gorm:"type:varchar(255);" json:"chapter"`
	ChatId  int64  `json:"chatId"`
}

func (b *Book) ToString() string {
	var res strings.Builder
	res.WriteString(strconv.Itoa(int(b.ID)))
	res.WriteString("\n")
	res.WriteString(b.Title)
	res.WriteString("\n")
	res.WriteString(b.Author)
	res.WriteString("\n")
	res.WriteString(b.Series)
	res.WriteString("\n")
	res.WriteString(b.Chapter)
	res.WriteString("\n")
	return res.String()
}

func (b *Book) Format() string {
	var res strings.Builder
	res.WriteString("(ID:" + strconv.Itoa(int(b.ID)) + ")")
	res.WriteString("\n<i><b>Название:</b></i> ")
	res.WriteString(b.Title)
	res.WriteString("\n<i><b>Автор:</b></i> ")
	res.WriteString(b.Author)
	res.WriteString("\n<i><b>Серия:</b></i> ")
	res.WriteString(b.Series)
	res.WriteString("\n<i><b>Глава:</b></i> ")
	res.WriteString(b.Chapter)
	res.WriteString("\nудалить: <code>/delete" + strconv.Itoa(int(b.ID)) + "</code>")
	res.WriteString("\n\n")
	return res.String()
}
