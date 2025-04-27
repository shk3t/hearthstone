package ui

import (
	"fmt"
	"hearthstone/internal/config"
	"hearthstone/internal/game"
	"hearthstone/pkg/sugar"
	"strings"
)

var DisplayFrame = sugar.If(config.Config.Debug, PrintFrame, UpdateFrame)

const prompt = "\n> "

func DisplayGame(topPlayer, botPlayer *game.Player, table *game.Table) {
	builder := strings.Builder{}
	fmt.Fprint(&builder, &topPlayer.Hand)
	fmt.Fprintln(&builder, &topPlayer.Hero)
	fmt.Fprint(&builder, table)
	fmt.Fprintln(&builder, &botPlayer.Hero)
	fmt.Fprint(&builder, &botPlayer.Hand)

	builder.WriteString(prompt)

	DisplayFrame(builder.String())
}

func UpdateFrame(content string) {
	fmt.Print("\033[2J\033[H")
	fmt.Print(content)
}

func PrintFrame(content string) {
	fmt.Print(content)
	fmt.Print("\n\n\n")
}