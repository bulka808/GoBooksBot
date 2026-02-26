package filter

import (
	S "GoGramTest/internal/state"

	tg "github.com/amarnathcjd/gogram/telegram"
)

type StateFilter struct {
	RequiredState int
	BotState      *S.BotState
}

func (s StateFilter) Check(m *tg.NewMessage) bool {
	if s.BotState.GetState() == s.RequiredState {
		return true
	}
	return false
}
func (s StateFilter) CheckCallback(c *tg.CallbackQuery) bool {
	if s.BotState.GetState() == s.RequiredState {
		return true
	}
	return false
}
func (s StateFilter) HasFlag(flag tg.FilterFlag) bool {
	return false
}

func NewStateFilter(bs *S.BotState, requiredState int) tg.Filter {
	return StateFilter{BotState: bs, RequiredState: requiredState}
}
