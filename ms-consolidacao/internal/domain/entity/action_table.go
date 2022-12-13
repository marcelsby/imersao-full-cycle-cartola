package entity

import "errors"

type ActionTableInterface interface {
	Init()
	GetScore(action string) (int, error)
}

type ActionTable struct {
	Action map[string]int
}

func (a *ActionTable) Init() {
	a.Action = map[string]int{
		"yellow card": -1,
		"red card":    -3,
		"goal":        5,
		"assist":      3,
	}
}

func (a *ActionTable) GetScore(action string) (int, error) {
	score, ok := a.Action[action]

	if !ok {
		return 0, errors.New("invalid action")
	}

	return score, nil
}
