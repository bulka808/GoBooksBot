package state

import (
	r "GoGramTest/internal/repository"
	"log"
	"sync"
)

const (
	Idle = iota
	Add
	Update
)

type BotState struct {
	mu        sync.Mutex
	State     int
	Repo      *r.BookRepository
	UpdateIDs chan int
	AddId     chan uint
}

func NewBotState(repo *r.BookRepository) *BotState {
	return &BotState{State: Idle, Repo: repo, UpdateIDs: make(chan int, 100), AddId: make(chan uint, 1)}
}
func (bs *BotState) SetState(state int) {
	bs.mu.Lock()
	bs.State = state
	bs.mu.Unlock()
	switch state {
	case 0:
		log.Println("[BOT STATE] idle")
	case 1:
		log.Println("[BOT STATE] add")
	case 2:
		log.Println("[BOT STATE] update")
	}

}
func (bs *BotState) GetState() int {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.State
}
func (bs *BotState) AddUpdate(id int) {
	bs.UpdateIDs <- id
}
