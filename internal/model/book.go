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
	var builder strings.Builder
	builder.WriteString("(ID:" + strconv.Itoa(int(b.ID)) + ")")
	builder.WriteString("\n<i><b>Название:</b></i> ")
	builder.WriteString(b.Title)
	builder.WriteString("\n<i><b>Автор:</b></i> ")
	builder.WriteString(b.Author)
	builder.WriteString("\n<i><b>Серия:</b></i> ")
	builder.WriteString(b.Series)
	builder.WriteString("\n<i><b>Глава:</b></i> ")
	builder.WriteString(b.Chapter)
	builder.WriteString("\n\n")
	return builder.String()
}
