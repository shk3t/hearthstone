package tui

type uiState struct {
	hint       string
	nextAction *nextPlayerAction
}

var state = uiState{hint: ""}

var nextActionHint = "Выберите цели"