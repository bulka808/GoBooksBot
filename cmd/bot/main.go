package main

import (
	"GoGramTest/internal/config"
	"GoGramTest/internal/database"
	f "GoGramTest/internal/filter"
	h "GoGramTest/internal/handler"
	"GoGramTest/internal/repository"
	s "GoGramTest/internal/state"
	"log"

	tg "github.com/amarnathcjd/gogram/telegram"
)

func main() {
	log.Println(tg.Version, tg.ApiVersion)
	//загрузка конфига
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// загрузка дб
	db, ctx := database.InitDB()
	//создание репы с книгами

	repo := repository.NewBookRepository(db, ctx)

	//создание клиента тг
	clientConfig := tg.NewClientConfigBuilder(int32(cfg.ApiId), cfg.ApiHash).
		WithSessionName("gogram").
		WithParseMode("HTML").
		Build()
	client, err := tg.NewClient(clientConfig)
	if err != nil {
		log.Fatal(err)
	}
	// подключение
	_, err = client.Conn()
	if err != nil {
		log.Fatal(err)
	}
	// вход
	_, err = client.Login(cfg.Phone)
	if err != nil {
		log.Fatal(err)
	}

	// создаем состояние бота
	botState := s.NewBotState(repo)

	// хэндлеры

	// по 2 хэндлера на команду
	// первый ловит команду, мб шлет сообшения, второй взаимодействует с ботом
	// add - добавление книги(ответом на команду для скачивания с бота)
	client.OnCommand("add", func(m *tg.NewMessage) error {
		return h.AddCommandHandler(m, botState)
	}).
		Filter(f.NewStateFilter(botState, s.Idle)).
		Filter(tg.FromUser(int64(cfg.OwnerID))).
		Filter(tg.Not(tg.IsBot)).
		Filter(tg.IsReply).
		Register()
	client.OnMessage("", func(m *tg.NewMessage) error {
		return h.AddBookHandler(m, botState)
	}).
		Filter(tg.FromUser(client.GetPeerID("botbybase_bot"))).
		Filter(f.NewStateFilter(botState, s.Add)).
		Filter(tg.FilterDocument).
		Register()
	// update - обновление и показ списка книг
	client.OnCommand("update", func(m *tg.NewMessage) error {
		return h.UpdateCommandHandler(m, botState)
	}).
		Filter(f.NewStateFilter(botState, s.Idle)).
		Filter(tg.FromUser(int64(cfg.OwnerID))).
		Filter(tg.Not(tg.IsBot)).
		Register()
	client.OnMessage("", func(m *tg.NewMessage) error {
		return h.UpdateBookHandler(m, botState)
	}).
		Filter(tg.FromUser(client.GetPeerID("botbybase_bot"))).
		Filter(f.NewStateFilter(botState, s.Update)).
		Filter(tg.FilterDocument).
		Register()
	// show - просто показ списка книг
	client.OnCommand("show", func(m *tg.NewMessage) error {
		return h.ShowCommandHandler(m, botState)
	}).
		Filter(f.NewStateFilter(botState, s.Idle)).
		Filter(tg.FromUser(int64(cfg.OwnerID))).
		Filter(tg.Not(tg.IsBot)).
		Register()
	// delete<id> уделение книги
	client.OnMessage("^/delete\\d+$", func(m *tg.NewMessage) error {
		return h.DeleteCommandHandler(m, botState)
	}).
		Filter(f.NewStateFilter(botState, s.Idle)).
		Filter(tg.FromUser(int64(cfg.OwnerID))).
		Filter(tg.Not(tg.IsBot)).
		Register()

	_, _ = client.SendMessage(cfg.OwnerID, "Ready!\n<code>/update</code>\n<code>/show</code>")

	client.Idle()
}
