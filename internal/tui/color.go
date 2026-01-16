package tui

import (
	"fmt"
	"hearthstone/internal/game"
	"hearthstone/pkg/ui"

	"github.com/fatih/color"
)

func getColorFunc(side game.Side) ui.FormatFunc {
	switch side {
	case game.TopSide:
		return color.CyanString
	case game.BotSide:
		return color.GreenString
	default:
		return fmt.Sprintf
	}
}
